package main

import (
	"fmt"
	"time"
)

func main() {
	ss := make(chan bool)
	go func() {
		tic := time.NewTicker(time.Second * 2)
		tic2 := time.NewTicker(time.Second * 10)
		for {
			select {
			case <-tic.C:
				ss <- true
			case <-tic2.C:
				ss <- false
			}
		}
	}()

	for {
		select {
		case t, v := <-ss:
			fmt.Println("t:", t, "v：", v)
		}
	}

}
