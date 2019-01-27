// Go in action
// @jeffotoni
// 2019-01-16

package main

func main() {

	j := 10
	i := 0
	switch j {
	case 11:
		println("here: 11")
		break
	default:
		println("here default")
		break
	}

	// infinitely
	for ; ; i++ {

		switch i {
		case 5:
			goto LABELS
		case i:
			println("i: ", i)
			break
		default:
			println("default: ", i)
		}
	}

LABELS:
	f()

}

func f() {
	println("goto fim")
}
