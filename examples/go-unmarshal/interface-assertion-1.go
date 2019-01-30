// Go in action
// @jeffotoni
// 2019-01-29

package main

import "fmt"

func doInterface(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	doInterface(33)
	doInterface("DevOpsBH")
	doInterface("Lambda")
	doInterface(true)
}
