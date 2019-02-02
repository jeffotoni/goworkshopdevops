// Go in action
// @jeffotoni
// 2019-01-01

package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	addr = ":8080"
)

// write bufio to optimization
func write(text string) {
	// var writer *bufio.Writer
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(text)
	writer.Flush()
}

func main() {

	// our function
	pingHandler := func(w http.ResponseWriter, req *http.Request) {
		json := `{"status":"success", "msg":"DevopsBH, Golang for Devops!"}`
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, json)
	}

	// HandleFunc
	http.HandleFunc("/v1/api/ping", pingHandler)

	// show
	write("\033[0;33mServer Run " 
	+ "Port " +
		addr + "\033[0m\n")

	// Listen
	log.Fatal(http.ListenAndServe(addr, nil))
}
