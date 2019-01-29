// Go in action
// @jeffotoni
// 2019-01-29

package main

import "fmt"
import "reflect"

func main() {
	data := []string{"one", "two", "three", "for", "five"}
	test(data)
	moredata := []int{1, 2, 3, 4, 5}
	test(moredata)
}

func test(t interface{}) {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)

		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
		}
	}
}
