package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type EntryType int32

const (
	EntryNomal EntryType = 0
	Entrydada  EntryType = 1
)

type Sss struct {
	Enrt EntryType
}

func main() {
	var j int
	tt := []int{2, 4, 6, 7, 8, 11, 14, 16, 19, 23, 24}
	for i := 0; i < 25; i += 3 {
		fmt.Println(i)
		if i == 0 {
			continue
		}

		if tt[j] <= i {
			//fmt.Println(tt[j])
		}

		j++
	}

	fmt.Println(CheckSign("ddc82e40f2183d280e5eeb29508550d9", "b0baae0630f444b0811ea3c2eb212170", "/app/v1/auth/autoLogin", "1615889127"))
}

func CheckSign(sign, appid, path, ts string) (ok bool) {
	hash := md5.New()
	hash.Write([]byte(appid))
	hash.Write([]byte(path))
	hash.Write([]byte(ts))
	calSign := hex.EncodeToString(hash.Sum(nil))
	return sign == calSign
}
