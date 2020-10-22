package main

import "fmt"

func main() {

	ch := make(chan string, 3)

	ch <- "func() {}"
	ch <- "func() {}"

	close(ch)
	//for value := range ch {
	//	fmt.Println("value:", value)
	//}
	fmt.Println("value:", <-ch)
	fmt.Println("value:", <-ch)
	fmt.Println("value:", <-ch)
	fmt.Println("value:", <-ch)

	ss := make(map[int]int, 1010)
	fmt.Println(ss, len(ss))
}
