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
	handlerfunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "DevopsBH for Golang simple one %s\n", r.URL.Path)
	})

	log.Printf("\nServer run 8080\n")
	err := http.ListenAndServe(":8080", handlerfunc)
	log.Fatal(err)
}
