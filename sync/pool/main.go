package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return "123"
		},
	}
	t := pool.Get().(string)
	fmt.Println(t)
	runtime.GC()
	//pool.Put("321")
	//runtime.GC()
	//pool.Put("222")
	//runtime.GC()
	//pool.Put("444")

	t2 := pool.Get().(string)
	fmt.Println(t2)
}
