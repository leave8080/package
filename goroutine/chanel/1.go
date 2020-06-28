package main

import "fmt"

func test(c chan struct{}) {
	d := <-c
	fmt.Println(d)
	close(c)
}
func main() {
	c := make(chan struct{})
	go test(c)
	c <- struct{}{}
	//d := <-c
	//fmt.Println(d)

}
