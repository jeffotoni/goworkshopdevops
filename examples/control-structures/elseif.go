// Go in action
// @jeffotoni
// 2019-01-16

package main

func main() {
	n := 100
	if n > 0 && n <= 55 {
		println("n > 0 or n <= 55")
	} else if n > 56 && n < 70 {
		println("n > 56 and n < 70")
	} else {

		if n >= 100 {
			println(" else here.. n > 100")
		} else {
			println(" else here.. n > 70")
		}
	}
}
