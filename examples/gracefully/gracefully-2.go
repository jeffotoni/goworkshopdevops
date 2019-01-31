// Go in action
// @jeffotoni
// 2019-01-29

package main

import "time"

type Worker struct {
	quit chan struct{}
}

func New() *Worker {
	return &Worker{
		quit: make(chan struct{}),
	}
}

func (w *Worker) Start() {
	println("starting")
	go w.work()
}

func (w *Worker) Stop() {
	println("stopping")
	close(w.quit)
	// faux pending work
	time.Sleep(5 * time.Second)
}

func (w *Worker) work() {
	for {
		select {
		case <-time.After(time.Second):
			println("doing work")
		case <-w.quit:
			return
		}
	}
}

func main() {
	worker := New()
	worker.Start()
	//gracefully.Timeout = 3 * time.Second
	//gracefully.Shutdown()
	worker.Stop()
	println("exiting")
}
