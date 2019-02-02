// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"fmt"
	"reflect"
)

func dump_interface_array(args interface{}) {

	// getting the interface
	val := reflect.ValueOf(args)

	// test the type
	if val.Kind() == reflect.Array {
		fmt.Println("len = ", val.Len())
		for i := 0; i < val.Len(); i++ {
			e := val.Index(i)
			switch e.Kind() {
			case reflect.Int:
				fmt.Printf("%v, ", e.Int())
			case reflect.Float32:
				fallthrough
			case reflect.Float64:
				fmt.Printf("%v, ", e.Float())
			case reflect.String:
				fmt.Printf("%v, ", e.String())
			default:
				panic(fmt.Sprintf("invalid Kind: %v", e.Kind()))
			}
		}
		fmt.Println()
	}
}

func main() {

	// array int
	int_ary := [4]int{1, 2, 3, 4}

	// array float
	float32_ary := [4]float32{1.1, 2.2, 3.3, 4.4}

	// array float
	float64_ary := [4]float64{1.1, 2.2, 3.3, 4.4}

	// array string
	string_ary := [...]string{"Scala", "Elixir",
		"Lisp", "Clojure"}

	dump_interface_array(int_ary)
	dump_interface_array(float32_ary)
	dump_interface_array(float64_ary)
	dump_interface_array(string_ary)
}
