/*
* Golang map generic
*
* @package     main
* @author      @jeffotoni
* @size        2018
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {

	interf := []int{1, 2, 3, 4, 5, 6, 7, 8}

	r := genericMap(interf, func(x int) string {
		return fmt.Sprintf("%d", x)
	})

	fmt.Println(r)

	interfs := []string{"jefferson", "Otoni", "Lima", "is", "A nice guy"}

	rs := genericMap(interfs, func(x string) string {
		return "" + x
	})

	fmt.Println(rs)
}

func genericMap(arr interface{}, mapFunc interface{}) interface{} {

	funcValue := reflect.ValueOf(mapFunc)
	arrValue := reflect.ValueOf(arr)

	arrType := arrValue.Type()
	arrElemType := arrType.Elem()

	if arrType.Kind() != reflect.Array && arrType.Kind() != reflect.Slice {
		panic("Parameter type is neither array nor slice.")
	}

	funcType := funcValue.Type()

	// Verificando se o segundo argumento é função
	// E também verificando se sua assinatura é func ({type A}) {type B}
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 1 || funcType.NumOut() != 1 {
		panic("Second argument must be map function")
	}

	if !arrElemType.ConvertibleTo(funcType.In(0)) {
		panic("Map function argument is not compatible with Array type")
	}

	resultSliceType := reflect.SliceOf(funcType.Out(0))

	resultSlice := reflect.MakeSlice(resultSliceType, 0, arrValue.Len())

	for i := 0; i < arrValue.Len(); i++ {

		resultSlice = reflect.Append(resultSlice, funcValue.Call([]reflect.Value{arrValue.Index(i)})[0])
	}

	// Convering resulting slice back to generic interface.
	return resultSlice.Interface()
}
