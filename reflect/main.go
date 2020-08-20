package main

import (
	"fmt"
	"reflect"
)

func main() {
	ss := "sss"
	num := 1

	fmt.Println(reflect.TypeOf(ss))

	fmt.Println(reflect.TypeOf(num))
}
