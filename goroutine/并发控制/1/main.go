package main

import (
	"fmt"
	"sync"
	"time"
)

func test(index int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("test=>", index)
}

func main() {
	maxchan := 1
	var wg sync.WaitGroup
	var ch = make(chan struct{}, maxchan)
	for i := 0; i < 100; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			defer func() {
				<-ch
			}()
			test(index)
		}(i)
	}
	wg.Wait()
}
