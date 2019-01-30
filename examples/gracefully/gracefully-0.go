// Go in action
// @jeffotoni
// 2019-01-24
// different ways to inizialize a struct

package main

import (
    "fmt"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

func worker(jobs <-chan int) {
    for job := range jobs {
        fmt.Println(job)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {

    jobs := make(chan int, 1000)

    var wg sync.WaitGroup

    // Spawn example workers
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(jobs <-chan int) {
            defer wg.Done()
            worker(jobs)
            fmt.Println("Worker is done.")
        }(jobs)
    }

    // Create example messages
    for i := 0; i < 100; i++ {
        jobs <- i
    }

    waitForSignal()

    close(jobs)
    fmt.Println("Closed jobs!")
    wg.Wait()
}

func waitForSignal() {
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, os.Interrupt)
    signal.Notify(sigs, syscall.SIGTERM)
    <-sigs
}
