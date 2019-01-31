// Go in action
// @jeffotoni
// 2019-01-31

package main

import (
	"fmt"
)

type User struct {
	FirstName string `tag_name:"firstname"`
	LastName  string `tag_name:"lastname"`
	Age       int    `tag_name:"age"`
}

func main() {

	// create instance
	// slice struct
	users := []*User{}

	user := new(User)
	user.FirstName = "Jefferson"
	user.LastName = "otoni"
	user.Age = 350
	users = append(users, user)

	user = new(User)
	user.FirstName = "Pike"
	user.LastName = "Hob"
	user.Age = 55
	users = append(users, user)

	fmt.Println(users)

	for k, v := range users {
		fmt.Println(k, v)
		fmt.Println(v.FirstName)
		fmt.Println(v.LastName)
		fmt.Println(v.Age)
	}
}
