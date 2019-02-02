// Go in action
// @jeffotoni
// 2019-01-01

package main

import (
	"fmt"
	"log"
	"net/http"
)

type pingHandler struct{}

func (h pingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DevopsBH for Golang simple %s\n", r.URL.Path)
}

func main() {
	log.Printf("\nServer run 8080\n")
	err := http.ListenAndServe(":8080", pingHandler{})
	log.Fatal(err)
}
