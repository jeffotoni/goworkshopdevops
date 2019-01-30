// Go in action
// @jeffotoni
// 2019-01-29

package main

import "fmt"

func main() {
	var i interface{}
	fmt.Println(i)

	i = "DevOps BH"
	fmt.Println(i)

	i = 2019
	fmt.Println(i)

	i = 9.5
	fmt.Println(i)

	i = [...]string{"Pike", "Jeffotoni", "Thompson", "Griesemer"}
	fmt.Println(i)

	i = map[string]interface{}{"Lang": []string{"Go", "Rust", "Scala", "Elixir"}}
	fmt.Println(i)

	//struct{ City string }{City: "BH"}
	i = map[struct{ L string }]interface{}{{L: "Lang"}: []string{"Go", "Rust", "Scala", "Elixir"}}
	fmt.Println(i)
}
