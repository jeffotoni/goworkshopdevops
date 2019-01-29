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
	Data string `json:"data"`
}

func main() {
	resp := `[{"data":"some data"}, {"data":"some more data"}]`

	data := &jsoninput{}
	_ = json.Unmarshal([]byte(resp), data)
	for _, value := range *data {
		fmt.Println(value.Data)
	}
}
