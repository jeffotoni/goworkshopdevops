// Go in action
// @jeffotoni
// 2019-01-24
// different ways to inizialize a struct

package main

import "fmt"

type Address struct {
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Zipcode      string `json:"zipcode"`
}

type ApiLogin struct {
	Name  string `json:"name"`
	Cpf   string `json:"cpf"`
	Login string `json:"login"`
	Email string `json:"email"`

	// anonymous
	And1 *struct {
		City string
	}

	// pointer
	And2 *Address

	// list Address
	And3 []Address
}

func main() {

	// different ways to inizialize a struct
	//
	//

	apilogin1 := &ApiLogin{Name: "@jeffotoni", Cpf: "093.393.334-34",
		And1: &struct{ City string }{City: "BH"}}

	fmt.Println(apilogin1)
	fmt.Println(apilogin1.Name)
	fmt.Println(apilogin1.And1)
	fmt.Println(apilogin1.And1.City)

	apilogin2 := &ApiLogin{Name: "@jeffotoni",
		Cpf:  "093.393.334-34",
		And2: &Address{City: "BH"}}

	fmt.Println(apilogin2)
	fmt.Println(apilogin2.Name)
	fmt.Println(apilogin2.And2)
	fmt.Println(apilogin2.And2.City)

	apilogin3 := &ApiLogin{Name: "@jeffotoni", Cpf: "093.393.334-34",
		And1: &struct{ City string }{City: "BH"},
		And2: &Address{City: "BH"}}

	fmt.Println(apilogin3)
	fmt.Println(apilogin3.Name)
	fmt.Println(apilogin3.And1)
	fmt.Println(apilogin3.And1.City)
	fmt.Println(apilogin3.And2)
	fmt.Println(apilogin3.And2.City)

	var apilogin4 ApiLogin
	fmt.Println(apilogin4)

	apilogin5 := ApiLogin{}
	fmt.Println(apilogin5)

	apilogin6 := &ApiLogin{}
	fmt.Println(apilogin6)

	apilogin7 := new(ApiLogin)
	fmt.Println(apilogin7)

	// another way to feed the struct
	g1add := Address{City: "Belo Horizonte"}
	g2add := Address{City: "Curitiba"}

	// declaring as list
	gall := []Address{}

	// add items
	gall = append(gall, g1add)
	gall = append(gall, g2add)

	fmt.Println(gall)

	// initializes Struct
	apil3 := ApiLogin{}

	// recive same type
	apil3.And3 = gall

	// show struct
	fmt.Println(apil3)

	// another way to initialize and feed the struct list
	apil3.And3 = []Address{{City: "Sao Paulo"}, {City: "Brasilia"}}

	// show struct
	fmt.Println(apil3)
}
