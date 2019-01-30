// Go in action
// @jeffotoni
// 2019-01-29

package main

import "fmt"

type I interface {
	comeon()
}

type A struct{}

func (a A) comeon() {}

type B struct{}

func (b B) comeon() {}

func main() {
	var i I
	i = A{} // dynamic type of i is A
	fmt.Printf("%T\n", i.(A))
	i = B{} // dynamic type of i is B
	fmt.Printf("%T\n", i.(B))
}
