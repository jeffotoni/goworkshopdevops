// Go in action
// @jeffotoni
// 2019-01-31

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	FirstName string `tag_name:"firstname"`
	LastName  string `tag_name:"lastname"`
	Age       int    `tag_name:"age"`
}

func (f *User) reflect() {

	// slice the element of struct
	val := reflect.ValueOf(f).Elem()

	// loop in elemnts of struct
	for i := 0; i < val.NumField(); i++ {

		// value of interface
		valueField := val.Field(i)

		// object of struct
		typeField := val.Type().Field(i)

		// tag of field
		tag := typeField.Tag
		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	}
}

func main() {

	f := &User{
		FirstName: "Jefferson",
		LastName:  "Otoni",
		Age:       350,
	}

	f.reflect()
}
