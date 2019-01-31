// Go in action
// @jeffotoni
// 2019-01-31

package main

import "fmt"

func PrintAll(vals []string) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	names := []string{"Rob Pike", "Jeffotoni", "Denis"}
	PrintAll(names)
}
