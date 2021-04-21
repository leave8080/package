package main

import (
	"fmt"
)

func main() {
	//fmt.Printf("%b\r\n", 8)                // 1000
	//fmt.Printf("%b\r\n", valueAtBit(8, 3)) // 0
	//fmt.Printf("%b\r\n", valueAtBit(8, 4)) // 1
	fmt.Println(valueAtBit(8, 8))
	//fmt.Printf("\n\n")
	//ascii()
	//stringFunc()
}

func valueAtBit(num, bit int) int {
	return (num >> (bit - 1)) & 1
}

func ascii() {
	var c rune = 'a'
	var i int = 98
	i1 := int(c)
	fmt.Println("'a' convert to", i1)
	c1 := rune(i)
	fmt.Println("98 convert to", string(c1))

	//string to rune
	for _, char := range []rune("世界你好") {
		fmt.Println(string(char))
	}

}

func stringFunc() {
	var ss = "abcdefghilmnopq"
	var sl []string
	ret := ""
	for i, v := range ss {
		ret += string(v)
		if (i+1)%5 == 0 {
			sl = append(sl, ret)
			ret = ""
		}
	}

	fmt.Println("ret", sl)

	for _, v := range sl {
		//j := string(v[0])
		k := string(v[1:3])
		//for _, v := range []rune(j) {
		//	fmt.Println(v)
		//}
		//fmt.Println(rune(v[0]))
		s := ""
		for i, v := range []rune(k) {
			fmt.Println(v)
			s += String(v)
			if i == 0 {
				s += ":"
			}
		}
		fmt.Printf(s)
		//fmt.Println([]rune(k))
		fmt.Printf("\n")
		//fmt.Println(string(v[0]), string(v[1:3]))

	}

}

func String(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
