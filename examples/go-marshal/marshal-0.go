// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
    "encoding/json"
    "fmt"
)

func main() {

    // let's try something simple
    // to understand what's
    // really going on
    m, _ := json.Marshal(false)
    fmt.Println(m)         // show as bytes
    fmt.Println(string(m)) // show as string

    m2, _ := json.Marshal(-1)
    fmt.Println(string(m2))

    m3, _ := json.Marshal(0)
    fmt.Println(string(m3))

    m4, _ := json.Marshal(102.3)
    fmt.Println(string(m4))

    m5, _ := json.Marshal("DevOpsBH")
    fmt.Println(string(m5))

    // For this to occur we need to pass a collection
    // with fields, something like a struct, a slice,
    // maps, interfaces {} Let's see below an example
    m6, _ := json.Marshal([...]string{"Elixir", "Scala"})
    fmt.Println(string(m6))

    m7, _ := json.Marshal(map[string]string{"twitter": "@jeffotoni"})
    fmt.Println(string(m7))

    m8, _ := json.Marshal(map[string]interface{}{"instagram": "jeffotoni", "langs": struct{ g string }{g: "gophers"}})
    fmt.Println(string(m8))

    m9, _ := json.Marshal(map[string]struct{ l string }{})
    fmt.Println(string(m9))
}
