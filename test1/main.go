package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				go test()
			}
		}()
	}
	wg.Wait()

}
func test() {
	fmt.Println("go")
}

//7， 设计一段代码限制goroutine的数量，比如有100个并行任务，将goroutine的数量限制在十个。
