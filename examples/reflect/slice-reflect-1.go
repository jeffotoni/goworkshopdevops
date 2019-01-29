// Go in action
// @jeffotoni
// 2019-01-29

package main

import "fmt"

func main() {
    loop([]string{"one", "two", "three"})
    loop([]int{1, 2, 3})
}

func loop(t interface{}) {
    switch t := t.(type) {
    case []string:
        for _, value := range t {
            fmt.Println(value)
        }
    case []int:
        for _, value := range t {
            fmt.Println(value)
        }
    }
}
