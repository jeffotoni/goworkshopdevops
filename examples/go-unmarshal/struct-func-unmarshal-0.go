// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Message struct {
	Subject  string `json:"subject"`
	TimeSent int64  `json:"sent"`
	Body     MyType `json:"body"`
	Status   string `json:"status"`
}

type MyType string

func (m *MyType) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch v.(type) {
	case int, float64, bool, map[string]interface{}:
		*m = MyType(b)
	case string:
		*m = MyType(b[1 : len(b)-1])
	default:
		return errors.New("unknown MyType")
	}
	return nil
}

func main() {
	m := &Message{}

	test1 := `{"subject":"test","sent":4567,"body":"body","status":"ok"}`
	if err := json.Unmarshal([]byte(test1), m); err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("m: %v\n", *m)

	test2 := `{"subject":"test","sent":4567,"body":3.14159625,"status":"ok"}`
	if err := json.Unmarshal([]byte(test2), m); err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("m: %v\n", *m)

	test3 := `{"subject":"test","sent":4567,"body":true,"status":"ok"}`
	if err := json.Unmarshal([]byte(test3), m); err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("m: %v\n", *m)

	test4 := `{"subject":"test","sent":4567,"body":{"some":"other","json":true},"status":"ok"}`
	if err := json.Unmarshal([]byte(test4), m); err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("m: %v\n", *m)
}
