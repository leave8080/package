package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 100
	xRefl := reflect.ValueOf(&x).Elem()
	xType := xRefl.Type()
	fmt.Printf("The type of x is %s.\n", xType)
}
