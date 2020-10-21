package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	age   int8
	name  string
	class int8
}
type B struct {
	age   int8
	class int8
	name  string
}

type part struct {
	a bool
	b int32
	c int8
	d int64
	e byte
}

func main() {
	var c int
	var d string
	var e int8
	var f int64
	var g *int
	var h unsafe.Pointer
	a := A{age: 10, name: "ll", class: 1}
	b := B{age: 10, name: "ll", class: 1}
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c), unsafe.Sizeof(d), unsafe.Sizeof(e),
		unsafe.Sizeof(f), unsafe.Sizeof(g), unsafe.Sizeof(h))

	part1 := part{}
	fmt.Printf("bool size: %d, align: %d, offset: %d\n", unsafe.Sizeof(part1.a), unsafe.Alignof(part1.a), unsafe.Offsetof(part1.a))
	fmt.Printf("int32 size: %d, align: %d, offset: %d\n", unsafe.Sizeof(part1.b), unsafe.Alignof(part1.b), unsafe.Offsetof(part1.b))
	fmt.Printf("int8 size: %d, align: %d, offset: %d\n", unsafe.Sizeof(part1.c), unsafe.Alignof(part1.c), unsafe.Offsetof(part1.c))
	fmt.Printf("int64 size: %d, align: %d, offset: %d\n", unsafe.Sizeof(part1.d), unsafe.Alignof(part1.d), unsafe.Offsetof(part1.d))
	fmt.Printf("byte size: %d, align: %d, offset: %d\n", unsafe.Sizeof(part1.e), unsafe.Alignof(part1.e), unsafe.Offsetof(part1.e))
	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part1), unsafe.Alignof(part1))
	fmt.Println(unsafe.Sizeof(struct {
		i8  int8
		i32 string
		i16 int16
	}{}),
		unsafe.Alignof(struct {
			i8  int8
			i32 string
			i16 int16
		}{}),
	)

}
