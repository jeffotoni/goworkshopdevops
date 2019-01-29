// Go in action
// @jeffotoni
// 2019-01-24
// different ways to inizialize a struct

package main

import "fmt"

type jsoninput []struct {
	Data string `json:"data"`
}

func main() {

	data := &jsoninput{{Data: "some data"}, {Data: "some more data"},
		{Data: "some more data"}}
	fmt.Println(data)

	// output:
	//  &[{some data} {some more data} {some more data}]
}
