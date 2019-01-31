// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
    "fmt"
    "reflect"
)

func main() {
    Schema := map[string]interface{}{}
    Schema["int"] = 10
    Schema["string"] = "this is a string and map with interface"
    Schema["bool"] = false

    val := reflect.ValueOf(Schema)
    fmt.Println("VALUE = ", val)
    fmt.Println("KIND = ", val.Kind())

    if val.Kind() == reflect.Map {
        for _, e := range val.MapKeys() {
            v := val.MapIndex(e)
            switch t := v.Interface().(type) {
            case int:
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
