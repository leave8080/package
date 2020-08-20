package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	f()

	fmt.Println("out")
	time.Sleep(9999 * time.Minute)
}

func f() {
	t1 := time.NewTicker(time.Duration(5) * time.Second)
	defer runtime.GC()

	defer t1.Stop()
	for {
		go func() {
			for {
				select {

				case <-t1.C:

					fmt.Println("ssssssssssssssssssssss")
					return
				default:
					fmt.Println("1")
					time.Sleep(1 * time.Second)
				}
			}
		}()

		time.Sleep(time.Second * 2)

	}

}
