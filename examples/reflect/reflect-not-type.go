// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"fmt"
	"io"
	"reflect"
)

func main() {
	var (
		a []string
		b []interface{}
		c []io.Writer
		d map[string]interface{}
	)
	fmt.Println(fill(&a), a) // pass
	fmt.Println(fill(&b), b) // pass
	fmt.Println(fill(&c), c) // fail
	fmt.Println(fill(&c), d) // fail
}

func fill(i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("non-pointer %v", v.Type())
	}
	// get the value that the pointer v points to.
	v = v.Elem()
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("can't fill non-slice value")
	}
	v.Set(reflect.MakeSlice(v.Type(), 3, 3))
	// validate the type of the slice. see below.
	if !canAssign(v.Index(0)) {
		return fmt.Errorf("can't assign string to slice elements")
	}
	for i, w := range []string{"foo", "bar", "baz"} {
		v.Index(i).Set(reflect.ValueOf(w))
	}
	return nil
}

// we accept strings, or empty interfaces.
func canAssign(v reflect.Value) bool {
	return v.Kind() == reflect.String || (v.Kind() == reflect.Interface && v.NumMethod() == 0)
}
