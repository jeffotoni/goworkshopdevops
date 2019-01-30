// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// byte with json
	var jsonBlob = []byte(`[
	{"Name": "Golang", "Level": "10"},
	{"Name": "Rust",    "Level": "7"}
]`)

	// struct that is our model
	type Lang struct {
		Name  string
		Level string
	}

	var langs []Lang

	// convert Json to Struct
	err := json.Unmarshal(jsonBlob, &langs)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("\n%+v", langs)
	fmt.Printf("\n%+v", langs[0].Name)
	fmt.Printf("\n%+v", langs[0].Level)
	fmt.Printf("\n%+v", langs[1].Name)
	fmt.Printf("\n%+v", langs[1].Level)
}
