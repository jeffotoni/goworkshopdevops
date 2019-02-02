// Go in action
// @jeffotoni
// 2019-01-01

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

// HTMLServer represents the web service that serves up HTML
type GoServerHttp struct {
	server *http.Server
	wg     sync.WaitGroup
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`
	<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Golang/DevOpsBH</title>

  <style>
  body {
  background-color: #424242;
  color: #F6F6F6;
  text-align: center;
  font-family: Helvetica, Arial, sans-serif;
  font-size: 20px;
  }
  h1, h2, h3 {
  margin: 0;
  line-height: 1.5;
  }
  .print-container {
  background-color: rgba(0, 0, 0, .3);
  padding: 15px;
  margin: 30px auto;
  width: 50%;
  border-radius: 4px;
  }
</style>

</head>
<body>
  <div class="print-container">
  <h1>{{ .Name }}</h1>
  <h2>Workshop Golang for DevOps!</h2>
  </div>
</body>
</html>
	`))
}

func Ping(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`{"status":"success","msg":"Devops BH for Golang StartServer!"}`))
}

func main() {

	// DefaultServeMux
	mux := http.NewServeMux()

	// POST handler /api/v1/ping
	handlerApiPing := http.HandlerFunc(Ping)

	// handler ping
	mux.Handle("/v1/api/ping", handlerApiPing)

	// templates/index html
	// if you want to activate this handler, the directory templates
	// where the html file is located must
	// be sent next to the binary to work, as it needs to parse the html
	// mux.HandleFunc("/", tpl.ShowHtml)

	// this handler implements the version
	// that does not need the html file
	mux.Handle("/", http.HandlerFunc(indexHandler))

	// Create the HTML Server
	ApiServer := GoServerHttp{
		server: &http.Server{
			Addr:           ":8080",
			Handler:        mux,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   20 * time.Second,
			MaxHeaderBytes: 1 << 25, //32Mb
		},
	}

	go func() {

		log.Printf("\nServer run :8080\n")
		// service connections
		if err := ApiServer.server.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	var errs = make(chan error, 2)

	go func() {
		// Setting up signal capturing
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		errs <- fmt.Errorf("Notify here: %s", <-c)

	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	//<-errs

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<-quit
	<-stop

	log.Println("Shutdown Server ...")
	// ... here is the code to close all
	// ...
	// ....

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ApiServer.server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown now:", err)
		// ... here is the code to close all error context
		// ...
		// ....
	}

	// execute finish
	log.Println("Server exist")

	<-errs
}
