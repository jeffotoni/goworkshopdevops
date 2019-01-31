// Go in action
// @jeffotoni
// 2019-01-31

package main

import "fmt"

func main() {
	names := []string{"Rob Pike", "Jeffotoni", "Denis"}
	for _, val := range names {
		fmt.Println(val)
	}
}
