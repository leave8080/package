package main

import (
	"fmt"
	"sync"
)

var num int
var m sync.Mutex
var s sync.Mutex

func main() {
	var ss sync.WaitGroup
	for i := 0; i < 5; i++ {
		ss.Add(1)
		go test(num, &ss)
	}
	for i := 0; i < 5; i++ {
		ss.Add(1)
		go test2(num, &ss)
	}
	ss.Wait()
}

func test(i int, ss *sync.WaitGroup) {
	defer ss.Done()
	i = add(i)
	fmt.Println("i:", i)
}
func test2(i int, ss *sync.WaitGroup) {
	defer ss.Done()
	i = del(i)
	fmt.Println("i:", i)
}
func add(i int) int {
	m.Lock()
	defer m.Unlock()
	num = i + 1
	return num
}
func del(i int) int {
	m.Lock()
	defer m.Unlock()
	num = i - 1
	return num
}

/*
package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex
	var arthmatic sync.WaitGroup

	Increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	Decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	for i := 0; i < 5; i++ {
		arthmatic.Add(1)
		go func() {
			defer arthmatic.Done()
			Increment()
		}()
	}

	for i := 0; i < 5; i++ {
		arthmatic.Add(1)
		go func() {
			defer arthmatic.Done()
			Decrement()
		}()
	}

	arthmatic.Wait()
	fmt.Printf("Arthmatic completed!\n")
}
*/
