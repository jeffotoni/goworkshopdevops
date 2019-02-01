// Go in action
// @jeffotoni
// 2019-01-01

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	// our function
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "DevopsBH, Golang for Devops!\n")
	}

	// handlerFunc
	http.HandleFunc("/v1/api/ping", helloHandler)

	// show
	log.Printf("Server Run :8080")

	// Listen
	log.Fatal(http.ListenAndServe(":8080", nil))
}
