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

type linkFetcher map[string]interface{}

type linkResult struct {
	body string
	urls []string
}

func main() {
	var l = linkFetcher{
		"https://golang.org/": linkResult{
			body: "The Go Programming Language",
			urls: []string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": linkResult{
			body: "Packages",
			urls: []string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		}}

	m, err := json.Marshal(l)

	if err != nil {
		log.Println(err)
	}
	fmt.Println("linkFetcher initialized")
	fmt.Println(l)

	fmt.Println("\njson.Marshal returning bytes")
	fmt.Println(m)

	fmt.Println("\njson.Marshal as string")
	fmt.Println(string(m))
}
