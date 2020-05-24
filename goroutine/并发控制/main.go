package main

import (
	"fmt"
	"github.com/leave8080/package/pool"
	"time"
)

func main() {
	t1 := time.Now()
	gopool := pool.NewGoPool(pool.WithMaxLimit(1))
	defer gopool.Wait()

	for i := 0; i < 10; i++ {
		gopool.Submit(func() {
			t := New(i, "d")
			fmt.Printf("%p", &t)
			fmt.Println(t)
		})
	}
	fmt.Println("@@@@@@")
	t2 := time.Now()
	time.Sleep(time.Second)
	fmt.Printf("%v", t2.Sub(t1))
}

var Test = New(2, "leave")

type RR struct {
	num  int
	Name string
}

func New(num int, name string) *RR {
	nn := &RR{
		num:  num,
		Name: name,
	}
	return nn
}
