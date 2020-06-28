package main

import "time"

func main() {
	i := make(chan int)
	c := make(chan string)
	go test(i, c)
	time.Sleep(3 * time.Second)
	i <- 1
	time.Sleep(5 * time.Second)
	c <- "c"
	time.Sleep(time.Second * 99)
}
func test(i chan int, c chan string) {
	for {
		select {
		case <-i:
			println("i")
		case <-c:
			println("c")
			//default:
			//	println("###")
		}
	}

}
