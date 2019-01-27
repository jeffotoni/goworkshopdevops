// Go in action
// @jeffotoni
// 2019-01-16

package main

func main() {

	i := 0
	// infinitely
	for ; ; i++ {
		for {

			if i == 10 {
				goto LABEL
			}
			i++
		}
	}

LABEL:
	f(i)

}

func f(i int) {
	println("label fim i: ", i)
}
