package main

import (
	"fmt"
	"reflect"
)

func main() {
	ss := "sss"
	num := 1
	t := reflect.TypeOf(ss)
	fmt.Println(t)

	fmt.Println(reflect.TypeOf(num))
	s := reflect.ValueOf(ss)
	fmt.Println(s)

	s = reflect.ValueOf(&s)
	t = reflect.TypeOf(&t)
	fmt.Println(s)
	fmt.Println(t)

}
