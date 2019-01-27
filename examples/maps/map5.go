// Go in action
// @jeffotoni
// 2019-01-24

package main

import "fmt"

func main() {
	var country = map[string]string{
		"Brasil": "Brasilia",
		"EUA":    "Washington, D.c",
		"Italy":  "Roma",
		"France": "Paris",
		"Japan":  "Toquio",
	}
	delete(country, "Japan")
	delete(country, "Italy")
	fmt.Println(country)
}
