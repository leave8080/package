package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	go pump(ch1) // pump hangs
	for {

		fmt.Println(<-ch1)

	}

	// prints only 1

}

func pump(ch chan int) {
	for i := 1; ; i++ {
		ch <- i
	}
}
