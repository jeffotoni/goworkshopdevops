// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	b := []byte(`{"Name":"Pike Hob","Age":56,"Parents":["Denis","James"]}`)
	var f interface{}
	_ = json.Unmarshal(b, &f)

	m := f.(map[string]interface{})

	// a way to initialize a map interface
	////////////////////////////////////////

	// f = map[string]interface{}{
	// 	"Name": "Wednesday",
	// 	"Age":  6,
	// 	"Parents": []interface{}{
	// 		"Torvalds",
	// 		"Mark Zuckerberg",
	// 	},
	// }
	/////////////////////////////////////

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "|is string|", vv)
		case float64:
			fmt.Println(k, "|is float64|", vv)
		case []interface{}:
			fmt.Println(k, "|is an array|")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
