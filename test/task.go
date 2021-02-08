package main

import "fmt"

type EntryType int32

const (
	EntryNomal EntryType = 0
	Entrydada  EntryType = 1
)

type Sss struct {
	Enrt EntryType
}

func main() {
	s := Sss{Enrt: Entrydada}
	key := "reee"
	fmt.Println("key \"" + key + "\" dadad ")
	fmt.Println(s)
}
