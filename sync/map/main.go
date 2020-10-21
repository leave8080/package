package main

import (
	"fmt"
	"sync"
)

var mp sync.Map

func main() {
	mp.Store("m1", "k1")
	mp.Store("m2", "k2")
	K1, ok := mp.Load("m1")
	if !ok {
		fmt.Println("err")
		return
	}
	fmt.Println(K1)
}
