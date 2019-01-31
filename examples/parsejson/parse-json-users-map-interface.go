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

	// Open our jsonFile
	jsonFile, err := os.Open("./user-only.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	//fmt.Println(result)
	obj := result["users"]
	//objData := obj.(map[string]interface{})
	objData := obj.([]interface{})

	//fmt.Println(objData)
	for k, v := range objData {

		fmt.Println("k:", k, "v:", v)
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
					fmt.Println("not found")
				}
			}
		}
	}

	// fmt.Println(objData["name"].(string))
	// fmt.Println(objData["type"].(string))
	// fmt.Println(objData["age"].(float64))

	//var v string
	//IUsers := result["users"]
	//fmt.Println(result["users"].(string))

	// for Key, value := range result {
	// 	fmt.Println("Key: ", Key, "Value:", value)

	// 	val := reflect.ValueOf(args)
	// 	fmt.Println(val.Kind())
	// }

	// fmt.Println("#######")
	// val := reflect.ValueOf(IUsers)
	// fmt.Println("Value: ", val)
	// fmt.Println("Kind: ", val.Kind())
	// fmt.Println("Type: ", reflect.TypeOf(IUsers).Elem())
}
