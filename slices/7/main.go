package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ss := "24123201021412320101"
	s := SplitSubN(ss, 10)
	for _, v := range s {
		fmt.Println(v)
		fmt.Println(string(v[2:4]))
		fmt.Println(Tool_DecimalByteSlice2HexString([]byte(v[2:4])))
		fmt.Println(Hex2Dec(v[4:6]))
	}
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

func Tool_DecimalByteSlice2HexString(DecimalSlice []byte) string {
	var sa = make([]string, 0)
	for _, v := range DecimalSlice {
		sa = append(sa, fmt.Sprintf("%02X", v))
	}
	ss := strings.Join(sa, "")
	return ss
}

func Hex2Dec(val string) int {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)
}
