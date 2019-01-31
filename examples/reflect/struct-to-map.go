// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

type address map[string]interface{}

type User struct {
	Login   string  `json:"login"`
	Email   string  `json:"email"`
	Status  string  `json:"status"`
	Address address `json:"address"`
}

func main() {

	a := address{
		"Anddress": address{
			"City": "Belo Horizonte",
			"Zip":  "34.566-333",
			"Fone": address{
				"fone1": "87-77047345",
				"fone2": "83-93483838",
			},
		},
	}

	user := User{
		Login:   "jeffotoni",
		Email:   "jeffotoni@gm.com",
		Status:  "alive",
		Address: a,
	}

	urlValues := structToMap(&user)
	fmt.Println(urlValues)
}

// convert struct to map
func structToMap(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		// You ca use tags here...
		// tag := typ.Field(i).Tag.Get("tagname")
		// Convert each type into a string for the url.Values string map
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()

		default:

			switch f.Kind() {
			// map
			case reflect.Map:
				for _, key := range f.MapKeys() {
					strct := f.MapIndex(key)
					//fmt.Println(key.Interface(), strct.Interface())
					v = fmt.Sprintf("%v %v", key.Interface(), strct.Interface())

				}
			}
		}

		values.Set(typ.Field(i).Name, v)
	}
	return
}
