// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// map more interface
	input := map[string]interface{}{
		"Key1": []string{"some", "key"},
		"key3": nil,
		"val":  2,
		"val2": "str",
		"val3": 4,
	}

	fmt.Println(input)
	for key, value := range input {
		slice, ok := value.([]string)
		if ok {
			input["Count"+key] = len(slice)
		}
	}
	fmt.Println(input)

	// Encode to Json
	m, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
	}

	// convert byte to string
	fmt.Println(string(m))

	// interface empty
	var f interface{}

	// Decode json
	if err := json.Unmarshal(m, &f); err != nil {
		fmt.Println(err)
	}

	// show screen
	fmt.Println(f)

	fmt.Println("### loop interface")
	m2 := f.(map[string]interface{})
	loopI(m2)
}

// loop in interface
func loopI(m2 map[string]interface{}) {
	for k, v := range m2 {

		// assert interface
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
}
