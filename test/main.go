package main

import (
	"fmt"
)

type Student struct {
	name  string
	code  int
	types int
}

func main() {

	s := make([]*Student, 5)
	s[0] = &Student{name: "qq",
		code:  5,
		types: 1,
	}
	s[1] = &Student{name: "qq",
		code:  2,
		types: 2,
	}
	s[2] = &Student{name: "aa",
		code:  2,
		types: 1,
	}
	s[3] = &Student{name: "aa",
		code:  1,
		types: 2,
	}
	s[4] = &Student{name: "sd",
		code:  2,
		types: 2,
	}
	fmt.Println(s)
	ret := make(map[string][]int)
	//for k,v := range s{
	//	fmt.Println(k,v.name,v.code,v.types)
	//
	//	for k,t := range s{
	//		if v.name == t.name{
	//			if v.types == 1{
	//				ret[v.name]=[]int{v.code,t.code}
	//			}else {
	//				ret[v.name]=[]int{t.code,v.code}
	//			}
	//
	//		}
	//	}
	//}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)-1; j++ {
			if s[i].name == s[j].name {
				fmt.Println(s[i])
				if s[i].types == 1 {
					ret[s[i].name] = []int{s[i].code, s[j].code}
				} else {
					ret[s[i].name] = []int{s[j].code, s[i].code}
				}
			} else {
				//ret[s[i].name]=[]int{s[i].code}
			}
		}
	}
	//fmt.Println(ret["qq"])
	//fmt.Println(ret["aa"])
	fmt.Println(ret)
}
