// Go in action
// @jeffotoni
// 2019-01-16

package main

import "fmt"

func main() {

	// string arrays / slice
	var lang = [...]string{"Erlang", "Elixir", "Haskell", "Clojure", "Scala"}

	// screen list
	fmt.Println(lang)

	// list the positions srtring arrays
	for k, v := range lang {
		fmt.Println(k, v)
	}

	/* create a map*/
	countryCapitalMap := map[string]string{"Brasil": "Brasilia", "EUA AMERICA": "Washington, D.C.", "France": "Paris", "Italy": "Roma", "Japan": "Tokyo"}

	/* print map using key-value*/
	for country, capital := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", capital)
	}

	// channel
	jobs := make(chan int, 3)

	// for channel
	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	// println(<-jobs)
	// println(<-jobs)
	// println(<-jobs)

	// close
	/* date is required for range to work*/
	close(jobs)

	/* This syntax is valid too. */
	for range jobs {
	}

	/* it is mandatory to close the channels to be able to scroll */
	for ch := range jobs {
		println(ch)
	}

	// it is not an array struct, it will range from error.
	sa := struct{ nick string }{"@jeffotoni"}
	fmt.Println(sa.nick)

	// here the range will be able to list all struct
	a := []struct{ nick string }{{"@devopsbr"}, {"@go_br"}, {"@awsbrasil"}, {"@go_br"}, {"@devopsbh"}}
	for i, v := range a {
		fmt.Println(i, v.nick)
	}

	// struct pointer
	var testdata *struct {
		a *[3]int
	}
	for i := range testdata.a {
		// testdata.a is never evaluated; len(testdata.a) is constant
		// i ranges from 0 to 2
		fmt.Println(i)
	}

	// new example interface and range
	var key string
	var val interface{} // element type of m is assignable to val
	m := map[string]int{"mon": 0, "tue": 1, "wed": 2, "thu": 3, "fri": 4, "sat": 5, "sun": 6}
	for key, val = range m {
		fmt.Println(key, val)
	}
}
