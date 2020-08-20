package main

import (
	"fmt"
	"time"
)

func main() {
	ss := make(chan int)
	go func() {
		tim := time.NewTicker(10 * time.Second)
		defer tim.Stop()
		select {
		case <-tim.C:

		case ret := <-ss:
			fmt.Println(ret)

		}
	}()
	ss <- 5
	for {
		time.Sleep(time.Second)
	}
}
