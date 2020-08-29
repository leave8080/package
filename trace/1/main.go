package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main1() {
	runtime.GOMAXPROCS(1)
	f, err := os.Create("trace.output")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		fmt.Println(err)
	}
	defer trace.Stop()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calcSum(i, &wg)
	}
	wg.Wait()
}
func calcSum(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum, n int64
	for ; n < 1000000; n++ {
		sum += n
	}
	fmt.Println(i, sum)
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
