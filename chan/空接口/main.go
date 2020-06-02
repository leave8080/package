package main

func main() {
	s := make(chan struct{})
	go test(s)
	s <- struct{}{}
	//<-s
	close(s)

	<-s

	//fmt.Println(<-s)
}

func test(ch chan struct{}) {
	<-ch
}
