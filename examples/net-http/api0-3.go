// Go in action
// @jeffotoni
// 2019-01-01

package main

import (
	"log"
	"net/http"
)

func main() {

	// our function
	pingHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("\nDevops BH for Golang Handle tree!"))
	}

	// Handle and recive http.HandlerFunc
	http.Handle("/v1/api/ping", http.HandlerFunc(pingHandler))
	http.Handle("/v1/api/ping2", http.HandlerFunc(pingHandler))
	http.Handle("/v1/api/ping3", http.HandlerFunc(pingHandler))
	// http.Handle("/v1/api/ping", pingHandler) // error

	// show run
	log.Printf("\nServer run 8080\n")

	// Listen
	// log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(pingHandler))) ok
	log.Fatal(http.ListenAndServe(":8080", nil))
}
