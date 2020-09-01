package main

import "fmt"

func main1() {
	c := make(chan int)
	fmt.Println(<-c)
	//直接读取空channel的死锁
	//当一个channel中没有数据，而直接读取时，会发生死锁：
}
