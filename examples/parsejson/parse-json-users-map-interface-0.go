// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("./user-only-0.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened users.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// reading archive content
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// defining the interface map that will receive
	// the file and the format is dynamic
	var result map[string]interface{}

	// let's now run our Unmarshal and convert to objects
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result)
	fmt.Println(result["name"].(string))
	fmt.Println(result["city"].(string))
	fmt.Println(result["age"].(float64))
}
