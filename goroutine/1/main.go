package main

import (
	"fmt"
	"sync"
)

func main1() {
	ch := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("chan->", i)
			ch <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-ch
		fmt.Println("main->", i)
	}
}

func main2() {
	var wg = sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

func main3() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go f(i, &wg)
	}
	wg.Wait()
}

func f(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	wg.Done()
}

//实现控制并发
func main4() {
	userCount := 4
	ch := make(chan bool, 3)
	for i := 0; i < userCount; i++ {
		ch <- true
		go Read1(ch, i)
		//time.Sleep(time.Second)
	}

	//	time.Sleep(time.Second)
}

func Read1(ch chan bool, i int) {
	fmt.Printf("go func: %d\n", i)
	<-ch
}

var wg = sync.WaitGroup{}

func main5() {
	userCount := 10
	for i := 0; i < userCount; i++ {
		wg.Add(1)
		go Read(i)
	}

	wg.Wait()
}

func Read(i int) {
	defer wg.Done()
	fmt.Printf("go func: %d\n", i)
}
