package main

import (
	"fmt"
	"sync"
	"time"
)

var i int
var mutex sync.Mutex

func main1() {

	go func() {
		for {
			mutex.Lock()
			i++
			fmt.Println("func1", i)
			time.Sleep(time.Second)
			mutex.Unlock()
		}
	}()
	go func() {
		for {
			mutex.Lock()
			i++
			fmt.Println("func2", i)
			time.Sleep(time.Second)
			mutex.Unlock()
		}

	}()
	for {
		//	i++
		mutex.Lock()
		fmt.Println("main", i)
		time.Sleep(time.Second)
		mutex.Unlock()
	}
}

var c = make(chan int)
var a string

func f() {
	a = "hello, world" // (1)
	fmt.Println("44")
	<-c // (2)
}
func main() {
	go f()
	fmt.Println("22")
	c <- 0 // (3)
	fmt.Println("33")
	print(a) // (4)
}
