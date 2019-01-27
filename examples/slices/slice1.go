// Go in action
// @jeffotoni
// 2019-01-24

package main

import "fmt"

func main() {
	// Creating a slice using a slice literal
	var s = []string{"@jeffotoni", "@awsbrasil", "@devopsbh", "@go_br"}

	// Short hand declaration
	t := []int{2, 4, 8, 16, 32, 64}

	fmt.Println("s = ", s, len(s))
	fmt.Println("t = ", t, len(t))
}
