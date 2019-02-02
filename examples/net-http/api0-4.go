// Go in action
// @jeffotoni
// 2019-01-01

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// our function
	pingHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("\nDevops BH for Golang mux HandleFunc!"))
	}

	// handleFunc
	mux.HandleFunc("/v1/api/ping", pingHandler) // ok

	mux.HandleFunc("/v1/api/ping2", http.HandlerFunc(pingHandler)) // ok

	mux.HandleFunc("/v1/api/ping3", pingHandler) // ok

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "You're lost, go home devopsBH!")
	})

	log.Printf("\nServer run 8080\n")
	// Listen
	log.Fatal(http.ListenAndServe(":8080", mux))
}
