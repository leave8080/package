package main

import (
	"fmt"
	"github.com/leave8080/package/pool"
)

func main() {
	gopool := pool.NewGoPool(pool.WithMaxLimit(3))
	defer gopool.Wait()

	for i := 0; i < 10; i++ {
		gopool.Submit(func() { fmt.Println("hhhh") })
	}

}
