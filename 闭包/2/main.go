package main

import "fmt"

func main() {

	a := A()
	for i := 0; i < 5; i++ {
		fmt.Println(a(i))
	}
}

func A() func(int) int {
	return func(i int) int {
		return i + 10
	}
}
