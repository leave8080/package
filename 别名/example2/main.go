package main

import (
	"log"
	"reflect"
)

type T2 struct{}

// T1 is deprecated, please use T2
type T1 = T2

func main() {
	var t T1
	f(t)
}

func f(t interface{}) {
	t, ok := t.(T1)
	if !ok {
		log.Fatal("t is not a T1 type")
	}
	// print main.T2
	println(reflect.TypeOf(t).String())
}
