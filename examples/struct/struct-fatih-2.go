// Go in action
// @jeffotoni
// 2019-01-24
// different ways to inizialize a struct

package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/structs"
)

type Server struct {
	Name        string `json:"name,omitempty"`
	ID          int
	Enabled     bool
	Users       []string // not exported
	http.Server          // embedded
}

func main() {

	// Create a new struct type:

	server := &Server{
		Name:    "gopher",
		ID:      12345678,
		Users:   []string{"jeffotoni", "pike", "dennis", "ken"},
		Enabled: true,
	}

	// struct
	fmt.Println(server)

	// create struct fatih
	s := structs.New(server)
	m := s.Map() // Get a map[string]interface{}
	fmt.Println(m)

	v := s.Values() // Get a []interface{}
	fmt.Println(v)

	f := s.Fields() // Get a []*Field
	fmt.Println(f)

	n := s.Names() // Get a []string
	fmt.Println(n)

	name := s.Field("Name") // Get a *Field based on the given field name

	// Get the underlying value,  value => "gopher"
	value := name.Value().(string)
	fmt.Println(value)

	tagValue := name.Tag("json")
	fmt.Println(tagValue)

	f1, ok := s.FieldOk("Name") // Get a *Field based on the given field name
	fmt.Println(f1, ok)

	n2 := s.Name() // Get the struct name
	fmt.Println(n2)

	h := s.HasZero() // Check if any field is uninitialized
	fmt.Println(h)
}
