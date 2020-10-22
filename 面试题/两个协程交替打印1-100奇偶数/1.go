package main

import (
	"fmt"
	"time"
)

func main() {
	p := make(chan int)
	go g1(p)
	go g2(p)
	time.Sleep(2 * time.Second)
}

func g1(p chan int) {
	for i := 1; i <= 100; i++ {
		p <- i
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}

func g2(p chan int) {
	for i := 1; i <= 100; i++ {
		<-p
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
