// Go in action
// @jeffotoni
// 2019-01-24

package main

import "fmt"

func main() {

	src := []string{"Erlang", "Elixir", "Haskell", "Clojure", "Scala"}
	dest := make([]string, 2)

	numElementsCopied := copy(dest, src)

	fmt.Println("src = ", src)
	fmt.Println("dest = ", dest)
	fmt.Println("Number of elements copied from src to dest = ", numElementsCopied)
}
