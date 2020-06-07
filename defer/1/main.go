package main

import "fmt"

func a() int {
	i := 0
	defer fmt.Println(i)
	i++
	return i
}

func main() {
	fmt.Println("a", a())
	fmt.Println(c())
	fmt.Println("###")
	t()
	t1()
}

func c() (i int) {
	fmt.Println(&i, "q")
	defer func() {
		i++
		fmt.Println(&i, "e")
	}()
	fmt.Println(&i, "r")
	return 1
}

//defer
func t() {
	a := make([]int, 0)
	defer fmt.Println("a =", a)
	a = append(a, 1)
}

//闭包引用传递
func t1() {
	a := make([]int, 0)
	defer func() { fmt.Println("a =", a) }()
	a = append(a, 1)
}
