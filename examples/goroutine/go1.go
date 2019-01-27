// Go in action
// @jeffotoni
// 2019-01-24

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.Intn(500)

func pinger() {
	time.Sleep(time.Duration(r) * time.Microsecond)
	fmt.Println("pinger")
}

func ponger() {
	time.Sleep(time.Duration(r) * time.Microsecond)
	fmt.Println("ponger")
}

func printer() {
	time.Sleep(time.Duration(r) * time.Microsecond)
	fmt.Println("printer")
}

func main() {
	// making functions
	// calls asynchronously
	go pinger()
	go ponger()
	go printer()

	// Waiting to press any key to end
	var input string
	fmt.Scanln(&input)
}
