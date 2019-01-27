// Go in action
// @jeffotoni
// 2019-01-24

package main

import (
	"encoding/json"
	"fmt"
)

var Error error
var b []byte

type ApiCpu struct {
	Server string
	Limit  int
	Spent  int
	Used   int
}

func main() {

	// Create an instance of the Box struct.
	// cpu := ApiCpu{
	// 	Server: "Apollo's",
	// 	Limit:  90,
	// 	Spent:  80,
	// 	Used:   10,
	// }
	fmt.Println(Error)
	cpu := make(chan int)
	// Create JSON from the instance data.
	b, Error = json.Marshal(cpu)
	if Error != nil {
		fmt.Println(Error)
	} else {
		fmt.Println(string(b))
	}
}
