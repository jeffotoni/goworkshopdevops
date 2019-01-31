// Go in action
// @jeffotoni
// 2019-01-24
// different ways to inizialize a struct

package main

import (
	"encoding/json"
	"fmt"
)

type jsoninput []struct {
	Name string `json:"name"`
}

func main() {

	// json in memory
	resp := `[{"name":"Andre"},{"name":"Pike"}]`

	// initialization struct
	data := &jsoninput{}

	// Unmarshal in bytes
	_ = json.Unmarshal([]byte(resp), data)

	// loop to see the values in the fields
	// loop in struct
	for _, value := range *data {
		fmt.Println(value.Name)
	}
}
