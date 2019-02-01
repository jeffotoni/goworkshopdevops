// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	// Open our jsonFile
	byteJson, err := ioutil.ReadFile("./user-only-0.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return
	}

	// show screen
	fmt.Println("Successfully Opened users.json")

	// defining the interface map that will receive
	// the file and the format is dynamic
	var result = make(map[string]interface{})

	// let's now run our Unmarshal and convert to objects
	json.Unmarshal(byteJson, &result)

	// show all
	fmt.Println(result)

	// how is a map I can do exactly
	// the syntax below
	val, ok := result["name"].(string)
	fmt.Println(val, ok)

	// type string
	fmt.Println(result["name"].(string))

	// type string
	fmt.Println(result["city"].(string))

	// type float64
	fmt.Println(result["age"].(float64))
}
