// Go in action
// @jeffotoni
// 2019-01-29

package main

import "fmt"

type T struct{}

func main() {

	s := T{}
	seen := make(map[struct{}]struct{}) // set of strings
	// ...
	if _, ok := seen[s]; !ok {
		seen[s] = struct{}{}
		// ...first time seeing s...
	}
	fmt.Println(seen)
}
