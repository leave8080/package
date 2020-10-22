package main

import (
	"fmt"
	"time"
)

func main() {
	s := make(chan int, 1)

	go func() {
		for {
			select {
			case v, ok := <-s:
				fmt.Println("<-s", v, ok)
			case s <- 1:
				fmt.Println("s<-", 1)
				//default:
				//	fmt.Println("fuck")
				//	time.Sleep(time.Second)
			}
		}
	}()
	for {
		time.Sleep(time.Hour)

	}
}
