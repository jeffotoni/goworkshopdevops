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

type Login struct {
	// Field appears in JSON as key "login".
	Login string `json:"login"`

	// Field appears in JSON as key "email" and
	// the field is omitted from the object if its value is empty,
	// as defined above.
	Email string `json:"email,omitempty"`

	// Field appears in JSON as key "nick" (the default), but
	// the field is skipped if empty.
	// Note the leading comma.
	Nick string `json:",omitempty"`

	// Field is ignored by this package.
	Level int `json:"-"`

	// Field appears in JSON as key "-".
	LastEmail string `json:"-,"`
}

func main() {

	l := Login{Login: "Austin", Email: "austin@go.com", Nick: "", Level: 1000, LastEmail: "austin@gmail.com"}
	fmt.Println(l)

	m, err := json.Marshal(l)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(m))
}
