package main

import (
	"fmt"
)

// Iterator 声明接口
type Iterator interface {
	Iterator(m iterSet) Iter
}

// Iter 迭代器的实现
type Iter struct {
	C chan interface{}
}

func newIter(i *iterSet) Iter {
	iter := Iter{make(chan interface{})}

	go func() {
		for k := range i.m {
			iter.C <- k
		}
		close(iter.C)
	}()

	return iter
}

// 我们自己的set
type iterSet struct {
	m map[string]bool
}

// Add 添加元素
func (i *iterSet) Add(k string) {
	i.m[k] = true
}

// Iterator 返回一个迭代器
func (i *iterSet) Iterator() Iter {
	return newIter(i)
}

func main() {
	aSet := iterSet{map[string]bool{}}
	aSet.Add("hello")
	aSet.Add("hello")
	aSet.Add("world")
	aSet.Add("world")

	iter := aSet.Iterator()

	for v := range iter.C {
		fmt.Printf("key: %s\n", v.(string))
	}
}
