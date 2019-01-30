// Go in action
// @jeffotoni
// 2019-01-24

package main

import (
	"fmt"
)

// We use the omitempty tag in 'json: field, omitempty'
// when we want the field not to appear after
// making the marshal if it is empty.
type D struct {
	Height int
	Width  int `json:"width"`
}

// Fields that do not have the "omitempty" tag will display
// the same field being empty when generating
// json through marshal.
type Cat struct {
	Breed    string `json:"breed,omitempty"`
	WeightKg int    `json:WeightKg,`
	Size     D      `json:"size,omitempty"`
}

func main() {
	d := Cat{
		//Breed:    "Persa",
		WeightKg: 23,
		Size:     D{20, 60},
	}
	//b, _ := json.Marshal(d)
	//fmt.Println(string(b))
	fmt.Println(d)
}
