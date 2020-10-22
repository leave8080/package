package main

import (
	"fmt"
)

type X struct {
	Name string
}
type Y struct {
	X
	Name string //相同名字的属性名会覆盖父类的属性
}

func main() {
	y := Y{X{"XChenys"}, "YChenys"}
	fmt.Println("y.Name = ", y.Name) //y.Name = YChenys
}

//所有的Y类型的Name成员的访问都只会访问到最外层的那个Name变量，X.Name变量相当于被覆盖了，可以用y.X.Name引用
