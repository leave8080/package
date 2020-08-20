package main

import (
	"fmt"
	"os"
	"sync"
)

var Once sync.Once

type Start struct {
	name string
}

func NewAdapter() *Start {
	fmt.Println("hhh", os.Getpid())
	return &Start{name: "leave"}
}

func main() {
	Once.Do(func() {
		NewAdapter()
		fmt.Println("first")
	})
	Once.Do(func() {
		NewAdapter()
		fmt.Println("second")
	})
}
