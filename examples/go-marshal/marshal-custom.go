// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"
)

type Email struct {
	Email string
}

type Time struct {
	time.Time
}

// layout
const layout = "2006-01-02"

// regex to email
var regxmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Time.Format(layout))), nil
}

// validate email
func (e Email) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, checkEmail(e.Email))), nil
}

func checkEmail(email string) string {
	if regxmail.MatchString(email) {
		return email
	} else {
		return "-invalid-"
	}
}

func main() {
	a := struct {
		Login string
		Email Email
		Time  Time
	}{
		Login: "devops",
		Email: Email{"jeffotoni-go.com"},
		Time:  Time{time.Now()},
	}

	fmt.Println(a)
	json, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(json))
}
