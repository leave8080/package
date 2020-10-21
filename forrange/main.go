package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhao1", Age: 12},
		{Name: "zhao2", Age: 13},
		{Name: "zhao3", Age: 14},
	}
	for i, stu := range stus {
		fmt.Println(&stu.Name)
		m[stu.Name] = &stu
	}
	for _, v := range m {
		fmt.Println(v.Name, v.Age)
	}

	fmt.Println("=====修改======")
	//for i := 0; i < len(stus); i++ {
	//	m[stus[i].Name] = &stus[i]
	//}
	for _, v := range stus {
		name := v
		fmt.Println(&name.Name)
		m[v.Name] = &name
	}
	for _, v := range m {
		fmt.Println(v.Name, v.Age)
	}
}
