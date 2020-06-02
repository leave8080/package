package main

import "fmt"

func main() {
	var v func(a int) int
	v = func(a int) int {
		return a + a
	}
	fmt.Println(v(1))
	t1 := add(2)
	fmt.Println(t1(1))
}

func add(base int) func(int) int {
	f := func(i int) int {
		fmt.Println("i", i)
		base += i
		return base
	}
	return f
}
