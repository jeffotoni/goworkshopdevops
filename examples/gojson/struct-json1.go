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

type Anddress struct {
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Zipcode      string `json:"zipcode"`
}

type ApiLogin struct {
	Name  string `json:"name"`
	Cpf   string `json:"cpf"`
	Login string `json:"login"`
	Email string `json:"email"`
	And1  *struct {
		City string
	}

	And2 *Anddress
}

func main() {

	apilogin1 := &ApiLogin{Name: "@jeffotoni", Cpf: "093.393.334-34", And1: &struct{ City string }{City: "BH"}, And2: &Anddress{City: "BH"}}
	m, err := json.Marshal(apilogin1)

	if err != nil {
		log.Println(err)
	}
	//fmt.Println("apilogin1 initialized")
	//fmt.Println(apilogin1)

	//fmt.Println("\njson.Marshal returning bytes")
	//fmt.Println(m)

	fmt.Println("\njson.Marshal as string")
	fmt.Println(string(m))
}
