// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
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

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	if string(b) == `null` {
		*t = Time{}
		return
	}
	t.Time, err = time.Parse(layout, string(b))
	return
}

// Unmarshal Json
func (t *Email) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	if string(b) == `null` {
		return
	}
	var stuff map[string]string
	err = json.Unmarshal(b, &stuff)
	if err != nil {
		return err
	}
	for key, value := range stuff {
		if strings.ToLower(key) == "email" {
			t.Email = checkEmail(value)
		}
	}
	return nil
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
	}{}

	b := []byte(`{"Login":"devops","Email":{"Email":"devops-go.com"},"Time":"2019-01-30"}`)
	err := json.Unmarshal(b, &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
