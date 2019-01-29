// Go in action
// @jeffotoni
// 2019-01-24
// different ways to inizialize a struct

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// creating our struct
// with some fields
// interesting as [] string
// and showing how to initialize them
type linkResult []struct {
	Body   string   `json:"body"`
	Urls   []string `json:"urls"`
	myname string   `json:"myname"`
}

func main() {

	var ll = linkResult{
		{
			Body:   "The Go Programming Language",
			Urls:   []string{"https://golang.org/pkg/", "https://golang.org/cmd/"},
			myname: "lambda man",
		},
		{
			Body:   "Package",
			Urls:   []string{"https://golang.org/", "https://golang.org/cmd/", "https://golang.org/pkg/fmt/"},
			myname: "go_br in action",
		}}

	m0, err := json.Marshal(ll)

	if err != nil {
		log.Println(err)
	}
	fmt.Println("linkFetcher initialized")
	fmt.Println(ll)

	fmt.Println("\njson.Marshal returning bytes")
	fmt.Println(m0)

	fmt.Println("\njson.Marshal as string")
	fmt.Println(string(m0))
}
