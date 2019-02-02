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
		w.Write([]byte("\nDevops BH for Golang HandleFunc!"))
	}

	// handleFunc
	http.HandleFunc("/v1/api/ping", pingHandler)
	http.HandleFunc("/v1/api/ping2", pingHandler)
	http.HandleFunc("/v1/api/ping3", pingHandler)
	http.HandleFunc("/v1/api/ping4", pingHandler)
	http.HandleFunc("/v1/api/ping5", pingHandler)
	http.HandleFunc("/v1/api/ping6", pingHandler)

	log.Printf("\nServer run 8080\n")
	// Listen
	log.Fatal(http.ListenAndServe(":8080", nil))
}
