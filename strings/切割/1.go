package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func main() {
	//ss := "weqqeqeqeqe"
	//sss := strings.Split(ss, "_")
	//fmt.Println(sss[0])
	//
	//sl := make([][]int, 2)
	//sl[0] = []int{1, 2}
	//sl[1] = []int{3, 4}
	//fmt.Println(sl)
	ss := "1412320104"
	//fmt.Println(rune(ss[1]))
	//feedPlanSlices := SplitSubN("1412320104", 10)
	//var hour, min string
	//
	//for _, v := range feedPlanSlices {
	//	fmt.Println(len(v))
	//	if string(v[9]) == "4" {
	//		hour = strconv.Itoa(Hex2Dec(v[2:4]))
	//		min = strconv.Itoa(Hex2Dec(v[4:6]))
	//	}
	//}
	//fmt.Println(hour, min)

	V1(ss)
	NewFk("")
}
func Hex2Dec(val string) int {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)
}
func SplitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}
func V1(plans string) {
	var sl []string
	ret := ""
	for i, v := range plans {
		var r rune = rune(v)
		// 真正可以输出字符
		fmt.Println(string(r))

		ret += string(v)
		if (i+1)%5 == 0 {
			sl = append(sl, ret)
			ret = ""
		}
	}

	for _, v := range sl {
		k := v[1:3]

		stime := ""
		for i, v := range []rune(k) {

			var r rune = rune(v)
			// 真正可以输出字符
			fmt.Println("sss:", string(r))
			//fmt.Println(v)
			stime += String(v)
			if i == 0 {
				stime += ":"
			}
		}
		fmt.Println("stime:", stime)
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

func NewFk(dd string) {

}
