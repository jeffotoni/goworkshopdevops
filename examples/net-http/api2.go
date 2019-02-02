// Go in action
// @jeffotoni
// 2019-01-01

package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	addr = ":8080"
)

// show log on screen
func logf(method, uri, nameHandle string, timeHandler time.Duration) {

	expre := "\033[5m%s \033[0;103m%s\033[0m \033[0;93m%s\033[0m \033[1;44m%s\033[0m"
	log.Printf(expre, method, uri, nameHandle, timeHandler)
}

// write bufio to optimization
func write(text string) {
	// var writer *bufio.Writer
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(text)
	writer.Flush()
}

func Ping(w http.ResponseWriter, r *http.Request) {

	// start time
	start := time.Now()

	if http.MethodPost == strings.ToUpper(r.Method) {

		json := `{"status":"success", "msg":"DevopsBH, Golang for Devops!"}`
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		//io.WriteString(w, json)
		w.Write([]byte(json))

	} else {

		json := `{"status":"error", "msg":"method not supported, only POST"}`
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusUnauthorized)
		//io.WriteString(w, json)
		//w.Write([]byte(json))
		http.Error(w, json, http.StatusUnauthorized)
	}

	logf(r.Method,
		r.RequestURI,
		"Ping",
		time.Since(start))
}

func main() {

	// handlerFunc
	http.HandleFunc("/v1/api/ping", Ping)

	// show
	write("\033[0;33mServer Run " +
		"Port " +
		addr + "\033[0m\n")

	// Listen
	log.Fatal(http.ListenAndServe(addr, nil))
}
