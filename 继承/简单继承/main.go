package main

import (
	"fmt"
)

type Father struct {
	MingZi string
}

func (this *Father) Say() string {
	return "大家好，我叫 " + this.MingZi
}

type Child struct {
	//	MingZi string
	Father
}

func main() {
	c := new(Child)
	c.MingZi = "小明"
	fmt.Println(c.Say())
}
