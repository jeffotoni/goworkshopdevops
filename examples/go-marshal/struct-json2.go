// Go in action
// @jeffotoni
// 2019-01-24
// different ways to inizialize a struct

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	type BouncedRecipients []struct {
		EmailAddress string `json:"emailAddress"`
		Action       string `json:"action"`
		Status       string `json:"status"`
		//DiagnosticCode string `json:"diagnosticCode"`
	}

	type Bounce struct {
		BounceType string `json:"bounceType"`
		//BounceSubType string    `json:"bounceSubType"`
		//Timestamp     time.Time `json:"timestamp"`
		BR BouncedRecipients
	}

	type JsonMessage struct {
		NotificationType string `json:"notificationType"`
		B                Bounce
		From             []string `json:"from"`
	}

	l := &JsonMessage{NotificationType: "38733773737xxxx",
		B: Bounce{BounceType: "bounce type",
			BR: BouncedRecipients{
				{"devops@g.com", "permanet", "error"}, {"lambdaman@g.com", "complaint", "success"}}},
		From: []string{"from1@m.com", "from2@gm.com"}}

	m, err := json.Marshal(l)

	if err != nil {
		log.Println(err)
	}

	//fmt.Println("\033[1;40mlinkFetcher initialized\033[0m")
	//fmt.Println(l)

	//fmt.Println("\n\033[1;42mjson.Marshal returning bytes\033[0m")
	//fmt.Println(m)

	fmt.Println("\n\033[1;41mjson.Marshal as string\033[0m")
	fmt.Println(string(m))
}
