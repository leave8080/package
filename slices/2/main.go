package main

import (
	"fmt"
	"reflect"
)

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5}

	s := data[2:4]
	s[0] += 100
	s[1] += 200

	fmt.Println(s)
	fmt.Println(data)
	fmt.Println(&s[0])
	fmt.Println(&data[2])
	fmt.Println(reflect.TypeOf(data))
}
