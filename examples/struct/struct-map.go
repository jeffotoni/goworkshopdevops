// Go in action
// @jeffotoni
// 2019-01-29

package main

import "fmt"

// Go in action
// @jeffotoni
// 2019-01-24
// different ways to inizialize a struct

type Before struct {
    m map[string]string
}

func contrivedAfter(b interface{}) interface{} {
    return struct {
        Before
        s []string
    }{b.(Before), []string{"new value", "value two"}}
}

func main() {
    b := Before{map[string]string{"some": "value"}}
    a := contrivedAfter(b)
    fmt.Println(a)
}
