// Go in action
// @jeffotoni
// 2019-01-16

package main

func main() {
	// will be looping infinitely
	// for {
	// }

	// will run only once and exit
	for {
		break
	}

	n := 5
	for n > 0 {
		n--
		println(n)
	}
	// Output:
	// 4
	// 3
	// 2
	// 1
	// 0

	// declaring i no and increasing i
	for i := 0; i < 5; i++ {
		println(i)
	}
	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4

	n = 5
	for i := 0; i < n; i++ {
		if i <= 2 {
			continue
		} else {
			println("i > 2 = ", i)
		}
	}

	// Output:
	// i > 2 =  3
	// i > 2 =  4

	n = 5
	for i := 0; i < n; i++ {
		if i == 2 {
			break
		} else {
			println("i: ", i)
		}
	}
	// Output:
	// i:  0
	// i:  1

	// infinitely
	for ; ; i++ {
		println("i: ", i)
	}

}
