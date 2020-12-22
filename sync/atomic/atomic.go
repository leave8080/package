package main

import (
	"fmt"
	"sync"
)

func main() {
	Func1()
}

func Func1() {
	aMap := sync.Map{}
	aMap.Store("1", "2")
	aMap.Store("3", "4")
	aMap.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %s, value %s\n", key, value)
		return true
	})
}
