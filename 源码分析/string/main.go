package main

import "unsafe"

func main() {
	ss := "dd"
	println(unsafe.Sizeof(ss))
}
