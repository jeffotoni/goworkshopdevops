// Go in action
// @jeffotoni
// 2019-01-01

// generating .key and .csr
// openssl req -nodes -newkey rsa:2048 -keyout server.key -out server.csr -subj "/C=BR/ST=Minas/L=Belo Horizonte/O=s3wf Ltd./OU=IT/CN=localhost"

// generating server .crt or .pem
// openssl x509 -req -sha256 -in server.csr -signkey server.key -out server.crt -days 365

package main

import (
	"io"
	"log"
	"net/http"
)

var (
	addr = ":443"
)

func main() {

	http.HandleFunc("/v1/api/ping", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "DevopsBH, Golang for Devops TLS!\n")
	})

	// show
	log.Printf("Server Run %s TLS / https://localhost%s", addr, addr)

	// conf listen TLS
	err := http.ListenAndServeTLS(addr, "server.crt", "server.key", nil)
	log.Fatal(err)
}

// curl --insecure -i -XGET https://localhost:8443/v1/api/ping
// curl -k -i -XGET https://localhost:8443/v1/api/ping
