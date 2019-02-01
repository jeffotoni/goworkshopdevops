// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Open our jsonFile
	byteJson, err := ioutil.ReadFile("./users.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(byteJson)
	fmt.Println(string(byteJson))
}
