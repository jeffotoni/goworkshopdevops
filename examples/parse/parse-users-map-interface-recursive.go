// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func main() {

	// Open our jsonFile
	byteJson, err := ioutil.ReadFile("./users.json")

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

	obj := result["users"]

	// get []interface
	objData := obj.([]interface{})

	// recursive
	DecodeMapVetInterface(objData)
}

func DecodeMapVetInterface(objData []interface{}) {
	//fmt.Println(objData)
	for line, v := range objData {

		fmt.Println("#### line: ", line)

		// ValueOf returns a new Value initialized
		// to the concrete value stored in the interface i.
		val := reflect.ValueOf(v)

		// A Kind represents the
		// specific kind of type
		// that a Type represents.
		if val.Kind() == reflect.Map {

			// Loop the map
			for _, key := range val.MapKeys() {

				fmt.Println("..............................................")

				// Index
				v := val.MapIndex(key)
				switch iv := v.Interface().(type) {

				// v.Interface().(type)
				// here test the types
				case int:
					fmt.Println(key, iv)
				case float64:
					fmt.Println(key, iv)
				case string:
					fmt.Println(key, iv)
				case bool:
					fmt.Println(key, iv)
				default:
					fmt.Println(key)
					DecodeMapInterface(iv)
				}
			}
		}
	}
}

func DecodeMapInterface(v interface{}) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Map {
		for _, e := range val.MapKeys() {
			v := val.MapIndex(e)
			switch t := v.Interface().(type) {
			case int:
				fmt.Println(e, t)
			case float64:
				fmt.Println(e, t)
			case string:
				fmt.Println(e, t)
			case bool:
				fmt.Println(e, t)
			default:
				fmt.Println("not found!")
			}
		}
	}
}
