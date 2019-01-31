package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	// open file json
	data, err := ioutil.ReadFile("./payload.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// make map interface
	payload := make(map[string]interface{})

	// Unmarshal data
	err = json.Unmarshal(data, &payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	// call recursive
	recursivMap(payload)
}

// recursive Map
func recursivMap(payload map[string]interface{}) {
	for k, v := range payload {
		//fmt.Printf("%v %T: %v\n", k, v, v)
		switch v.(type) {
		case []interface{}:
			recursivSlice(v.([]interface{}))
		case map[string]interface{}:
			//fmt.Printf("%v %T: %v\n", k, v, v)
			fmt.Println("----------------------")
			fmt.Printf("%v\n", k)
			recursivMap(v.(map[string]interface{}))

		default:
			fmt.Printf("%v=%v\n", k, v)
		}
	}
}

// recursive Slice
func recursivSlice(pauload []interface{}) {
	for _, v := range pauload {
		switch v.(type) {
		case []interface{}:
			recursivSlice(v.([]interface{}))
		case map[string]interface{}:
			recursivMap(v.(map[string]interface{}))
		}
	}
}
