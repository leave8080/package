package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup
	f := func() {
		fmt.Println("!!!!!!!!")
	}
	for i := 0; i < 100; i++ {
		go func(index int) {
			wg.Add(1)
			defer wg.Done()
			once.Do(f) //多次调用只执行一次
		}(i)
	}
	wg.Wait()
}
