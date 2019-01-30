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

type ApiLogin struct {
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}

func main() {

	a := ApiLogin{"Jefferson", "033.343.434-89"}
	fmt.Println(a)

	m, err := json.Marshal(a)
	if err != nil {
		log.Println(err)
	}
	// show bytes
	fmt.Println(m)

	// show string json
	fmt.Println(string(m))

	// improving output for json format viewing
	json, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
}
