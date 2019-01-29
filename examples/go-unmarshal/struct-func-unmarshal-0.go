package main

import (
	"encoding/json"
	"errors"
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

// func main() {
// 	m := &Message{}

// 	test1 := `{"subject":"test","sent":4567,"body":"body","status":"ok"}`
// 	if err := json.Unmarshal([]byte(test1), m); err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Printf("m: %v\n", *m)

// 	test2 := `{"subject":"test","sent":4567,"body":3.14159625,"status":"ok"}`
// 	if err := json.Unmarshal([]byte(test2), m); err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Printf("m: %v\n", *m)

// 	test3 := `{"subject":"test","sent":4567,"body":true,"status":"ok"}`
// 	if err := json.Unmarshal([]byte(test3), m); err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Printf("m: %v\n", *m)

// 	test4 := `{"subject":"test","sent":4567,"body":{"some":"other","json":true},"status":"ok"}`
// 	if err := json.Unmarshal([]byte(test4), m); err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Printf("m: %v\n", *m)
// }
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Message struct {
// 	Subject  string `json:"subject"`
// 	TimeSent int64  `json:"sent"`
// 	Body     BodyMap `json:"body"`
// 	Status   string `json:"status"`
// }

// type BodyMap map[string]Datum

// type Datum interface{}

// type MyStringType string
// type MyNumType float64
// type MyBoolType bool

// func (m*BodyMap) UnmarshalJSON(b []byte) error {
// 	var myMap map[string]Datum
// 	if err := json.Unmarshal(b, &myMap); err != nil {
// 		return err
// 	}

// 	bm := make(map[string]Datum)
// 	for k,v := range myMap{
// 		switch v.(type) {
// 		case float64:
// 			bm[k] = MyNumType(v.(float64))
// 		case bool:
// 		      bm[k] = MyBoolType(v.(bool))
// 		case string:
// 			bm[k] = MyStringType(string(b)[1 : len(b)-1])
// 		default:
// 			return fmt.Errorf("unknown type")
// 		}
// 	}
// 	*m = bm
// 	return nil
// }

// func main() {
// 	m := &Message{}

// 	test1 := `{"subject":"test","sent":4567,"body":{"strField":"string","numField":4567,"boolField":true},"status":"ok"}`
// 	if err := json.Unmarshal([]byte(test1), m); err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Printf("m: %v\n", *m)
// }

// package main

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// )

// type Message struct {
// 	Subject  string `json:"subject"`
// 	TimeSent int64  `json:"sent"`
// 	Body     MyType `json:"body"`
// 	Status   string `json:"status"`
// }

// type MyType struct {
// 	Type string
// 	Val  string
// 	Src  []byte
// }

// func (m *MyType) UnmarshalJSON(b []byte) error {
// 	var v interface{}
// 	if err := json.Unmarshal(b, &v); err != nil {
// 		return err
// 	}
// 	switch v.(type) {
// 	case int:
// 		m.Type = "int"
// 		m.Val = string(b)
// 		m.Src = b
// 	case float64:
// 		m.Type = "float64"
// 		m.Val = string(b)
// 		m.Src = b
// 	case bool:
// 		m.Type = "bool"
// 		m.Val = string(b)
// 		m.Src = b
// 	case map[string]interface{}:
// 		m.Type = "map"
// 		m.Val = string(b)
// 		m.Src = b
// 	case string:
// 		m.Type = "string"
// 		m.Val = string(b[1 : len(b)-1])
// 		m.Src = b
// 	default:
// 		return errors.New("unknown MyType")
// 	}
// 	return nil
// }

// func main() {
// 	m := &Message{}

// 	test1 := `{"subject":"test","sent":4567,"body":"body","status":"ok"}`
// 	if err := json.Unmarshal([]byte(test1), m); err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Printf("m: %v\n", *m)

// 	test2 := `{"subject":"test","sent":4567,"body":3.14159625,"status":"ok"}`
// 	if err := json.Unmarshal([]byte(test2), m); err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Printf("m: %v\n", *m)

// 	test3 := `{"subject":"test","sent":4567,"body":true,"status":"ok"}`
// 	if err := json.Unmarshal([]byte(test3), m); err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Printf("m: %v\n", *m)

// 	test4 := `{"subject":"test","sent":4567,"body":{"some":"other","json":true},"status":"ok"}`
// 	if err := json.Unmarshal([]byte(test4), m); err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Printf("m: %v\n", *m)
// }
