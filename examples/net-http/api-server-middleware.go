// Go in action
// @jeffotoni
// 2019-01-01

package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

// color terminal
var Expre = "\033[5m%s \033[0;103m%s\033[0m \033[0;93m%s\033[0m \033[1;44m%s\033[0m"

func Ping(w http.ResponseWriter, r *http.Request) {

	json := `{"status":"success","msg":"pong"}`
	w.Write([]byte(json))
}

// This middleware is responsible for holding up when we have a
// very large number of accesses in a very small time interval,
// depending on the capacity of your traffic, cpu and memory.
// It is one of the favorite middlewares, it is very powerful,
// not only determines the number of clients, but it does
// not have to lose in the requests sent.
func MaxClients(n int) Adapter {
	sema := make(chan struct{}, n)
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sema <- struct{}{}
			defer func() { <-sema }()
			h.ServeHTTP(w, r)
		})
	}
}

// This middleware is only a simulation, to implement the
// jwt in Go is very quiet, I will
// demonstrate in other topics below.
func AuthJwt() Adapter {
	//s1 := logg.Start()
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//if gjwt.CheckJwt(w, r) {
			if r.Header.Get("X-KEY") == "123" {
				h.ServeHTTP(w, r)
			} else {
				msgjson := `{"status":"error","message":"error in Jwt!"}`
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusUnauthorized)
				io.WriteString(w, msgjson)
				//logg.Show(r.URL.Path, strings.ToUpper(r.Method), "error", s1)
			}
		})
	}
}

// Middleware Logger
func Logger(name string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			h.ServeHTTP(w, r)
			log.Printf(
				"%s %s %s %s",
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
			)
		})
	}
}

// Middleware Logger
func LoggerColor(name string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			h.ServeHTTP(w, r)
			log.Printf(
				Expre,
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
			)
		})
	}
}

type Adapter func(http.Handler) http.Handler

// Middleware
func Middleware(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func main() {

	mux := http.NewServeMux()

	handlerApiPing := http.HandlerFunc(Ping)

	// generate token jwt
	// handler token
	mux.Handle("/v1/api/ping",
		Middleware(handlerApiPing,
			Logger("ping"),
		))

	mux.Handle("/v1/api/login",
		Middleware(handlerApiPing,
			AuthJwt(),
			//Logger("login"),
			LoggerColor("login"),
		))

	// show run server
	log.Printf("\nServer run :8080\n")

	// Listen
	log.Fatal(http.ListenAndServe(":8080", mux))
}
