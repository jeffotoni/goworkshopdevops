// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	b := []byte(`{"Name":"DevOps","Age":10,"Company":["Google","Aws"]}`)

	var f interface{}
	if err := json.Unmarshal(b, &f); err != nil {
		fmt.Println(err)
	}

	fmt.Println(f)

	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	//fmt.Println()
}
