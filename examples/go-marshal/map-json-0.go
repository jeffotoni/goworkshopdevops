// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/** version struct
Anddress struct {
	City string `json:"City"`
	Fone struct {
		Fone1 string `json:"fone1"`
		Fone2 string `json:"fone2"`
	} `json:"Fone"`
	Zip string `json:"Zip"`
} `json:"Anddress"`
Data struct {
	Create bool   `json:"create"`
	Level  int    `json:"level"`
	Login  string `json:"login"`
} `json:"Data"`
*/

func main() {

	// empty interface
	// interface, to create dynamic structures
	type login map[string]interface{}

	a := login{
		"Data": login{
			"login":  "jeffotoni",
			"level":  1000,
			"create": true,
		}, "Anddress": login{
			"City": "Belo Horizonte",
			"Zip":  "34.566-333",
			"Fone": login{
				"fone1": "87-77047345",
				"fone2": "83-93483838",
			},
		},
	}

	s, _ := json.Marshal(a)
	fmt.Println(string(s))

	// improving output for json format viewing
	json, err := json.MarshalIndent(j, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
}
