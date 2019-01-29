// Go in action
// @jeffotoni
// 2019-01-24

package main

import "fmt"

type linkResult struct {
	body string
	urls []string
}

type linkFetcher map[string]*linkResult

func main() {
	// Required to initialize
	// the map with values
	var m1 map[string]int
	var m2 = make(map[string]int)
	var m3 = map[string]int{"population": 500000}
	var m4 = m3
	var m5 map[string]string
	/* create a map*/
	m5 = make(map[string]string)
	fmt.Println(m1, m2, m3, m4, m5)

	var l = linkFetcher{
		"https://golang.org/": &linkResult{
			"The Go Programming Language",
			[]string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": &linkResult{
			"Packages",
			[]string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		}}

	fmt.Println(l)
}
