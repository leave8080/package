package main

import "fmt"

//先的声明再初始化最后赋值
func main1() {
	var ss map[int]int
	fmt.Printf("%v\n", ss == nil)
	ss = make(map[int]int)
	fmt.Printf("%v\n", ss == nil)
	ss[1] = 1
	fmt.Println(ss)
}

//直接make进行初始化之后再赋值
func main2() {
	ss := make(map[int]int)
	ss[1] = 1
	fmt.Println(ss)
}

//直接初始化赋值
func main() {
	ss := map[int]int{
		1: 1,
		2: 2,
	}
	fmt.Println(ss)
}
