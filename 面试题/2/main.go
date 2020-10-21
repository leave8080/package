package main

import "fmt"

func A() func() int {

	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func main() {
	f := A()
	for i := 0; i < 10; i++ {
		fmt.Println(i, f())
	}
}
