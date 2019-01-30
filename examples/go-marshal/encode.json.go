// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
    "bytes"
    "encoding/json"
    "fmt"
)

const (
    empty = ""
    tab   = "\t"
)

func main() {

    //json, err := MyJson(map[string]string{"twitter": "@jeffotoni"})
    json, err := MyJson(map[string]interface{}{"instagram": "jeffotoni", "langs": struct{ G string }{G: "gophers"}})

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(json)
}

func MyJson(data interface{}) (string, error) {

    buffer := new(bytes.Buffer)
    encoder := json.NewEncoder(buffer)
    encoder.SetIndent(empty, tab)
    // encode...
    err := encoder.Encode(data)
    if err != nil {
        return empty, err
    }
    return buffer.String(), nil
}
