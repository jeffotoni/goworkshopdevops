// Go in action
// Code in Go vs C
// @jeffotoni
// 2019-01-24

package main

import (
        "fmt"
)

type User struct {
        Login string
}

type Users struct {
        Logins []User
}

func (u *Users) AddUser(login User) []User {
        u.Logins = append(u.Logins, login)
        return u.Logins
}

func main() {

        // add user list
        login1 := User{Login: "jeffotoni.lambdaman"}
        login2 := User{Login: "andre.almar"}

        // instance User
        listUser := []User{}
        u := Users{listUser}

        // add login
        u.AddUser(login1)
        u.AddUser(login2)

        // show struct
        fmt.Println(len(u.Logins))
        fmt.Println(u.Logins)
}
