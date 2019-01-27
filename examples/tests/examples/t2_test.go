// Go in action
// @jeffotoni
// 2019-01-24

package main

import "testing"

func TestSum(t *testing.T) {
	x := 1 + 1
	if x != 2 {
		t.Error("Error! 1 + 1 is not equal to 2, I got", x)
	}
}
