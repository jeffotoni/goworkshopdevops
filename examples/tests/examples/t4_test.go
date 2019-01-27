// Go in action
// @jeffotoni
// 2019-01-24

package main

import (
	"testing"

	"github.com/jeffotoni/goworkshopdevops/examples/tests/pkg/math"
)

func TestSum(t *testing.T) {
	x := 2
	y := 3
	want := 5

	t.Run("t4_test", func(t *testing.T) {
		if got := math.Sum(x, y); got != want {
			t.Errorf("Sum() = %v, want %v", got, want)
		}
	})
}
