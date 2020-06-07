package main

import "fmt"

func test1(x int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var a [10]int
	a[x] = 14
}
func main() {
	test1(12)
}
