package main

func worker(ch chan struct{}) {
	// Receive a message from the main program.
	println("3")
	<-ch
	println("roger2")

	// Send a message to the main program.
	close(ch)
}

func main() {
	ch := make(chan struct{})
	go worker(ch)
	println("1")
	// Send a message to a worker.
	//select {}

	ch <- struct{}{}
	println("2")
	// Receive a message from the worker.
	//ch <- struct{}{}
	//close(ch)
	<-ch
	println("roger")
	// Output:
	// roger
	// roger
}
