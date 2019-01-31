// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func main() {

	// We can use the ioutil.ReadFile which would already
	// return bytes and we would not need to use two functions
	// ioutil.ReadFile("./user-only.json")

	// Open our jsonFile
	jsonFile, err := os.Open("./user-only.json")

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

	// the syntax below is also allowed
	//var result = make(map[string]interface{})

	// let's now run our Unmarshal and convert to objects
	json.Unmarshal([]byte(byteValue), &result)

	// Let's access the interface on our map
	obj := result["users"]

	// Here completely changes from the
	// previous example we now have an [] interface,
	// ie an array of interfaces a slice.
	objInterface := obj.([]interface{})

	// Now as we have multiple users,
	// we'll loop through to fetch them
	// and access each user's key
	// and value in the slice
	for line, keyOne := range objInterface {

		fmt.Println("#### line ", line)

		// ValueOf returns a new Value initialized
		// to the concrete value stored in the interface i.
		val := reflect.ValueOf(keyOne)

		// A Kind represents the
		// specific kind of type
		// that a Type represents.
		if val.Kind() == reflect.Map {

			// Loop the map
			for _, key := range val.MapKeys() {

				// Index of Map
				v := val.MapIndex(key)

				// get the value of type Interface
				switch value := v.Interface().(type) {

				// v.Interface().(type)
				// here test the types
				case int:
					fmt.Println(key, value)
				case float64:
					fmt.Println(key, value)
				case string:
					fmt.Println(key, value)
				case bool:
					fmt.Println(key, value)
				default:
					fmt.Println("not found")
				}
			}
		}
	}
}
