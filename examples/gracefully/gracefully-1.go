package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// create a "returnCode" channel which will be the return code of the application
	var returnCode = make(chan int)

	// finishUP channel signals the application to finish up
	var finishUP = make(chan struct{})

	// done channel signals the signal handler that the application has completed
	var done = make(chan struct{})

	// gracefulStop is a channel of os.Signals that we will watch for -SIGTERM
	var gracefulStop = make(chan os.Signal)

	// watch for SIGTERM and SIGINT from the operating system, and notify the app on
	// the gracefulStop channel
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	// launch a worker whose job it is to always watch for gracefulStop signals
	go func() {
		// wait for our os signal to stop the app
		// on the graceful stop channel
		// this goroutine will block until we get an OS signal
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)

		// send message on "finish up" channel to tell the app to
		// gracefully shutdown
		finishUP <- struct{}{}

		// wait for word back if we finished or not
		select {
		case <-time.After(30 * time.Second):
			// timeout after 30 seconds waiting for app to finish,
			// our application should Exit(1)
			returnCode <- 1
		case <-done:
			// if we got a message on done, we finished, so end app
			// our application should Exit(0)
			returnCode <- 0
		}
	}()

	// ... Do business Logic in goroutines

	fmt.Println("waiting for finish")
	// wait for finishUP channel write to close the app down
	<-finishUP
	fmt.Println("stopping things, might take 2 seconds")

	// ... Do business Logic for shutdown simulated by Sleep 2 seconds
	time.Sleep(2 * time.Second)

	// write to the done channel to signal we are done.
	done <- struct{}{}
	os.Exit(<-returnCode)
}
