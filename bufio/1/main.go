package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.NewTicker(2 * time.Second)
	ch := make(chan bool)
	defer tick.Stop()
	go func() {
		for {
			select {
			case <-ch:
				return
			case <-tick.C:
				fmt.Println("2s")
			}

		}
	}()

	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
	time.Sleep(10 * time.Second)
	ch <- true

}
