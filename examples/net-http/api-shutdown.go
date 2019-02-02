// Go in action
// @jeffotoni
// 2019-01-01

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {

	pingHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("\nDevops BH for Golang Handle shutdown!"))
	}

	var errs = make(chan error, 2)

	//server := &http.Server{Addr: ":8080", Handler: http.FileServer(http.Dir("/."))}
	server := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(pingHandler)}

	go func() {
		errs <- server.ListenAndServe()
	}()

	go func() {
		// Setting up signal capturing
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		errs <- fmt.Errorf("%s", <-c)

	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	<-errs

}
