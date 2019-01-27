// Go in action
// @jeffotoni
// 2019-01-24

package main

import (
	"bufio"
	"os"
)

// creating the write object pointer
// so that we can receive value in every
// scope of our program
var writer *bufio.Writer

func main() {
	// All screen output will be redirected
	// to bufio.NewWriter
	writer = bufio.NewWriter(os.Stdout)
	s := "How many stars does Orion have?\n"
	var b byte = 'H'

	writer.WriteString(s)
	writer.WriteByte(b)
	writer.WriteString("\n")

	// when all the functions finishes it closes
	// the buffer and sends to the.Stdout
	defer writer.Flush()
}
