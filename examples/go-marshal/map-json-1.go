// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"encoding/json"
	"fmt"
)

// create type map
type body map[string]interface{}

// struct Message
type Message struct {
	Subject  string                 `json:"subject"`
	TimeSent int64                  `json:"sent"`
	Body     body                   `json:"body"`  // embedded => Body of the type map[string]interface{}
	Body2    map[string]interface{} `json:"body2"` // Body of the type map[string]interface{}
	Status   string                 `json:"status"`
}

func main() {

	// our interface map body
	// This map can be dynamic,
	// create types as needed, because the interface{}
	// is for this purpose, it creates dynamic types.
	b := body{
		"Data": body{
			"login":  "jeffotoni",
			"level":  1000,
			"create": true,
		}, "Anddress": body{
			"City": "Jersey City",
			"Zip":  "34.566-333",
			"Fone": body{
				"fone1": "87-77047345",
				"fone2": "83-93483838",
			},
		},
	}

	// create type
	type out map[string]interface{}

	// a way to initialize the struct
	// and the map contained within the struct
	s := &Message{
		Subject:  "Test Struct to Map",
		TimeSent: 345,
		Body:     b,
		Body2:    out{"Instance": "c5.xlarge", "vCpu": "4", "Ecu": "16", "Mem": "8"},
		Status:   "success",
	}
	fmt.Println(s)

	// Another way to implement is to generate
	// the map interface {} inside the initialization
	// of the struture.
	s2 := &Message{
		Subject:  "Test Struct to Map",
		TimeSent: 345,
		Body:     b,
		Body2:    map[string]interface{}{"Instance": "c5.xlarge", "vCpu": "4", "Ecu": "16", "Mem": "8"},
		Status:   "success",
	}
	fmt.Println(s2)

	// converting everything to Json
	m, err := json.Marshal(s2)
	if err != nil {
		fmt.Println(err)
	}

	// sending json to the screen
	fmt.Println("##### json")
	fmt.Println(string(m))
}
