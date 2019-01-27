// Go in action
// @jeffotoni
// 2019-01-16

package main

import "fmt"

func main() {
	n := 0

LOOP1:
	n++
	if n == 10 {
		println("fim")
		return
	}

	if n%2 == 0 {
		goto LOOP2
	} else {

		fmt.Println("n", n, "LOOP1 here...")
		goto LOOP1
	}

LOOP2:
	fmt.Println("n", n, "LOOP2 here...")
	goto LOOP1

}
