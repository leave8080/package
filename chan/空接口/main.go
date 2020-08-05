package main

import (
	"fmt"
	"time"
)

func main() {
	s := make(chan struct{})
	go test(s)
	//s <- struct{}{}
	//<-s
	time.Sleep(time.Second * 6)
	close(s)

	//	<-s

	//fmt.Println(<-s)
}

func test(ch chan struct{}) {
	for {
		select {
		case <-ch:
			return
		default:
			for {
				fmt.Println("1")
				time.Sleep(time.Second)
			}

		}
	}
	///	<-ch
}
