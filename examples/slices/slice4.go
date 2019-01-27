// Go in action
// @jeffotoni
// 2019-01-24

package main

import "fmt"

func main() {
	slice1 := []string{"Clojure", "Scala", "Elixir"}
	slice2 := append(slice1, "Assembly", "Rust", "Go")
	fmt.Printf("slice1 = %v, len = %d, cap = %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 = %v, len = %d, cap = %d\n", slice2, len(slice2), cap(slice2))
	slice1[0] = "C++"
	fmt.Println("\nslice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)

	var s []string

	// Appending to a nil slice
	s = append(s, "Java", "C", "Lisp", "Haskell")
	fmt.Printf("\ns = %v, len = %d, cap = %d\n", s, len(s), cap(s))

}
