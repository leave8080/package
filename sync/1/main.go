package main

func main() {
	c := make(chan int)
	c <- 88
	<-c
}
