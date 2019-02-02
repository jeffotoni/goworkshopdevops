// Go in action
// @jeffotoni
// 2019-01-01

package main

import (
	"fmt"
	"log"
	"net/http"
)

type numberDumperString string
type numberDumperInt int

// http HandlerFunc
func (n numberDumperString) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DevOps BH, Golang is Life, Here's your number: %s\n", n)
}

// http HandlerFunc
func (n numberDumperInt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DevOps BH, Golang is Life, Here's your number: %d\n", n)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/one", numberDumperString("one"))
	mux.Handle("/two", numberDumperString("two"))
	mux.Handle("/three", numberDumperInt(3))
	mux.Handle("/four", numberDumperInt(4))
	mux.Handle("/five", numberDumperInt(5))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "That's not a supported number new Ghoper!")
	})

	// show run
	log.Printf("\nServer run :8080\n")

	// listen
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
