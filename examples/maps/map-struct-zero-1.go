// Go in action
// @jeffotoni
// 2019-01-29

package main

import "fmt"

func main() {

	s := "key"
	seen := make(map[string]struct{}) // set of strings
	// ...
	if _, ok := seen[s]; !ok {
		seen[s] = struct{}{}
		// ...first time seeing s...
	}
	fmt.Println(seen)
}
