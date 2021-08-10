package main

import (
	"fmt"
	"github.com/remeh/sizedwaitgroup"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(math.MaxInt32)
	swg := sizedwaitgroup.New(5)
	for i := 0; i < 10; i++ {
		swg.Add()
		go func(i int) {
			defer swg.Done()
			query(i)
		}(i)
	}

	swg.Wait()
}
func query(i int) {
	fmt.Println(i)
	ms := i + 500 + rand.Intn(500)
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
