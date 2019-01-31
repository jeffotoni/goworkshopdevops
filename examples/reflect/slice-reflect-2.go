// Go in action
// @jeffotoni
// 2019-01-29

package main

import "fmt"
import "reflect"

func main() {

	// slice dinamic
	slice := []string{"C", "C++", "Fortram", "Cobol"}
	dump_slice(slice)

	// slice int
	dataint := []int{1, 2, 3}
	dump_slice(dataint)
}

// dump interfaces
func dump_slice(t interface{}) {

	// type kind only slice
	switch reflect.TypeOf(t).Kind() {

	// slice
	case reflect.Slice:

		// return interface
		s := reflect.ValueOf(t)

		// loop in type
		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
		}
	}
}
