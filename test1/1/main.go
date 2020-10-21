package main

import (
	"fmt"
	"time"
)

func main() {
	//var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		//	wg.Add(1)
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
}
