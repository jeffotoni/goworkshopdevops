// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"encoding/pem"
	"log"
	"os"
)

func main() {
	block := &pem.Block{
		Type: "MESSAGE",
		Headers: map[string]string{
			"Devops": "Gopher",
		},
		Bytes: []byte("your key here, test #$"),
	}

	// func Encode(out io.Writer, b *Block) error
	if err := pem.Encode(os.Stdout, block); err != nil {
		log.Fatal(err)
	}
}
