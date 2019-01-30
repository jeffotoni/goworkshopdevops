// Go in action
// @jeffotoni
// 2019-01-29

package main

import "fmt"

type Login struct {
	User, Login, Email string
}

// passing a struct as parameter
// for our struct map
var m = map[string]Login{
	"jeffotoni": {"jeffotoni", "jeff", "jeff@gm.com"},
	"Google":    {"root", "super", "google@gm.com"},
}

func main() {
	fmt.Println(m)
}
