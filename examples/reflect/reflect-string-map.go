// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
    "fmt"
    "reflect"
)

func main() {

    // map interface
    Schema := map[string]interface{}{}

    // let's fill in the fields
    Schema["int"] = 10
    Schema["string"] = "this is a string and map with interface"
    Schema["bool"] = false
    Schema["sliceString"] = [...]string{"C", "C++", "Lisp"}

    // interface
    val := reflect.ValueOf(Schema)

    // all data
    fmt.Println("Schema", val)

    // type
    fmt.Println("Type = ", val.Kind())

    // kind type == reflect map
    if val.Kind() == reflect.Map {
        for _, key := range val.MapKeys() {
            v := val.MapIndex(key)

            // kind type, interface type
            switch value := v.Interface().(type) {

            // testing types
            case int:
                fmt.Println(key, value)
            case string:
                fmt.Println(key, value)
            case bool:
                fmt.Println(key, value)

            // case map[string]string:

            // case []interface{}:

            // case map[string]interface{}:

            default:
                val2 := reflect.ValueOf(Schema["sliceString"])
                // type array
                if val2.Kind() == reflect.Array {
                    //fmt.Println("len = ", val2.Len())

                    // loop array
                    for i := 0; i < val2.Len(); i++ {
                        // index key
                        e := val2.Index(i)

                        // kind type
                        switch e.Kind() {

                        // case types
                        case reflect.String:
                            fmt.Printf("%v, ", e.String())

                        //default
                        default:
                            panic(fmt.Sprintf("invalid Kind: %v", e.Kind()))
                        }
                    }
                    fmt.Println()
                }
            }
        }
    }
}
