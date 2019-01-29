// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
    "fmt"
    "reflect"
)

func main() {

    key := 1
    value := "C++"

    mapType := reflect.MapOf(reflect.TypeOf(key), reflect.TypeOf(value))

    mapValue := reflect.MakeMap(mapType)
    mapValue.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
    mapValue.SetMapIndex(reflect.ValueOf(2), reflect.ValueOf("jeff"))
    mapValue.SetMapIndex(reflect.ValueOf(3), reflect.ValueOf("golng"))

    keys := mapValue.MapKeys()
    fmt.Println(keys)
    fmt.Println(mapValue)
    for _, k := range keys {
        c_key := k.Convert(mapValue.Type().Key())
        c_value := mapValue.MapIndex(c_key)
        fmt.Println("key :", c_key, " value:", c_value)
    }

}
