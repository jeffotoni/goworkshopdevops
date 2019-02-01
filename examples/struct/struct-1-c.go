// Go in action
// Code in Go vs C
// @jeffotoni
// 2019-01-24

package main

import "fmt"

// Go in action
// @jeffotoni
// 2019-01-24

// #include <stdio.h>
// #include <stdlib.h>

// struct Login {
//     char *User;
//     char *Email;
// };

//In Go would look like this
type Login struct {
    User  string
    Email string
}

// int main() {

// In Go
func main() {

    //struct Login login, *pointerToLogin;

    // In Go
    login := Login{User: "jeffotoni", Email: "jef@m.com"}

    login.User = "jeffotoni"
    login.Email = "jef@m.com"

    // pointerToLogin = malloc(sizeof(struct Login));
    // pointerToLogin->User = "pike";
    // pointerToLogin->Email = "pike@g.com";

    var pointerToLogin = new(Login)
    pointerToLogin.User = "pike"
    pointerToLogin.Email = "pike@g.com"

    //printf("login vals: %s %s\n", login.User, login.Email);
    //printf("pointerToLogin: %p %s %s\n", pointerToLogin, pointerToLogin->User, pointerToLogin->Email);

    fmt.Printf("login vals: %s %s\n", login.User, login.Email)
    fmt.Printf("pointerToLogin: %v %s %s\n", pointerToLogin, pointerToLogin.User, pointerToLogin.Email)

    //free(pointerToLogin);
    return
}
