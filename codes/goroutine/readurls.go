package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

/**
"as":"AS5769 Videotron Telecom Ltee",
"city":"Vaudreuil-Dorion",
"country":"Canada",
"countryCode":"CA",
"isp":"Le Groupe Videotron Ltee",
"lat":45.4,
"lon":-73.9333,
"org":"Videotron Ltee",
"query":"24.48.0.1",
"region":"QC",
"regionName":"Quebec",
"status":"success",
"timezone":"America/Toronto",
"zip":"H9X"
*/

type IpApi struct {
	Name     string `json:name`
	Url      string `json:url`
	City     string `json:city`
	Country  string `json:country`
	Log      string `json:log`
	Timezone string `json:timezone`
	TimeDiff string `json:timediff`
}

type Result string

func DoRequest(name, url string) Result {
	var contents []byte
	var err error
	t1 := time.Now()
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		return ""
	} else {
		defer response.Body.Close()
		contents, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			return ""
		}
	}

	var jsonIp = IpApi{}
	err = json.Unmarshal(contents, &jsonIp)
	if err != nil {
		fmt.Printf("%s", err)
		return ""
	}

	jsonIp.Name = name
	jsonIp.Url = url
	t2 := time.Now()
	diff := t2.Sub(t1)

	jsonIp.TimeDiff = fmt.Sprintf("%s", diff)
	b, err := json.Marshal(jsonIp)
	if err != nil {
		fmt.Printf("%s", err)
		return ""
	}

	// http.Get...
	//t := time.Duration(rand.Intn(150)) * time.Millisecond
	//time.Sleep(t)
	//return Result(fmt.Sprintf("%s in %s: %s: %s", name, url, string(contents), diff))
	return Result(fmt.Sprintf("%s", string(b)))
}

type Services map[string]string

func DoAllRequests(s Services) <-chan Result {
	c := make(chan Result, len(s))
	for i, v := range s {
		go func(k, v string) { c <- DoRequest(k, v) }(i, v)
	}
	return c
}

func main() {
	fmt.Println("#### start ###")
	rand.Seed(time.Now().UnixNano())
	srv := Services{
		"ip-api.com/fields":     "http://ip-api.com/json/24.48.0.1?fields=61439",
		"ip-api.com/json-en":    "http://ip-api.com/json/52.45.11.179",
		"ip-api.com/json/pt-br": "http://ip-api.com/json/104.109.21.88?lang=pt-BR",
	}
	results := DoAllRequests(srv)
	for i := 0; i < len(srv); i++ {
		select {
		case result := <-results:
			// json ...
			fmt.Println("#######")
			fmt.Println(result)
		case <-time.After(time.Millisecond * 150):
			// prefiro esse tipo de timeout em vez do ctx, mas vai de cada problema
			jsonRes := <-results
			var jsonIp = IpApi{}
			err := json.Unmarshal([]byte(jsonRes), &jsonIp)
			if err != nil {
				fmt.Printf("%s", err)
			}
			fmt.Println("timed out: [name]:", jsonIp.Name, "[Url]:", jsonIp.Url)
			//return
		}
	}
	// sem waitgroup :)
	return
}
