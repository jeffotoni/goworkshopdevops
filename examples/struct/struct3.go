// Go in action
// @jeffotoni
// 2019-01-24
// different ways to inizialize a struct

package main

import "fmt"

type Anddress struct {
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Zipcode      string `json:"zipcode"`
}

type ApiLogin struct {
	Name  string `json:"name"`
	Cpf   string `json:"cpf"`
	Login string `json:"login"`
	Email string `json:"email"`
	And1  *struct {
		City string
	}

	And2 *Anddress
}

func main() {

	// different ways to inizialize a struct
	//
	//

	apilogin1 := &ApiLogin{Name: "@jeffotoni", Cpf: "093.393.334-34", And1: &struct{ City string }{City: "BH"}}
	fmt.Println(apilogin1)
	fmt.Println(apilogin1.Name)
	fmt.Println(apilogin1.And1)
	fmt.Println(apilogin1.And1.City)

	apilogin2 := &ApiLogin{Name: "@jeffotoni", Cpf: "093.393.334-34", And2: &Anddress{City: "BH"}}
	fmt.Println(apilogin2)
	fmt.Println(apilogin2.Name)
	fmt.Println(apilogin2.And2)
	fmt.Println(apilogin2.And2.City)

	apilogin3 := &ApiLogin{Name: "@jeffotoni", Cpf: "093.393.334-34", And1: &struct{ City string }{City: "BH"}, And2: &Anddress{City: "BH"}}
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
}
