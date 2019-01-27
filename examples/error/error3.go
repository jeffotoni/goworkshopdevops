// Go in action
// @jeffotoni
// 2019-01-24

package main

import (
	"errors"
	"fmt"
	"math"
)

// By convention, errors are the last return value and
// have type `error`, a built-in interface.
func Sqrt(fvalue float64) (float64, error) {
	if fvalue < 0 {

		// `errors.New` constructs a basic `error` value
		// with the given error message.
		return 0, errors.New("Math: negative number passed to Sqrt [" + fmt.Sprintf("%.2f", fvalue) + "]")
	}

	// retunr sqrt
	return math.Sqrt(fvalue), nil
}

func main() {
	result, err := Sqrt(-33)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sqrt (-33) = [", result, "]")
	}

	result, err = Sqrt(81)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sqrt(81) = [", result, "]")
	}
}
