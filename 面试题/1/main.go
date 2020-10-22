package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func page() map[string]*Student {
	stu := []Student{
		{Name: "q",
			Age: 10,
		},
		{Name: "w",
			Age: 20,
		}, {Name: "e",
			Age: 30,
		},
	}
	ss := make(map[string]*Student, 3)
	//for _, v := range stu {
	//	ss[v.Name] = &v
	//	fmt.Println(&ss)
	//}
	for i, _ := range stu {
		ss[stu[i].Name] = &stu[i]
	}
	return ss
}
func main() {
	ss := page()
	for k, v := range ss {
		fmt.Println(k, v)
	}
}
