package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

// MyTime 自定义时间类型
type MyTime time.Time

// MarshalJSON MyTime 的序列化方法，转化为时间戳
func (t MyTime) MarshalJSON() ([]byte, error) {
	timeStramp := time.Time(t).Unix()
	return []byte(strconv.Itoa(int(timeStramp))), nil
}

// UnmarshalJSON 时间戳[]byte 转化为MyTime类型
func (t *MyTime) UnmarshalJSON(data []byte) error {
	timeStramp, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	*t = MyTime(time.Unix(int64(timeStramp), 0))
	return nil
}

// T asdf
type T struct {
	CreateTime MyTime `json:"create_time"`
}

func main() {
	tString, err := json.Marshal(T{CreateTime: MyTime(time.Now())})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("序列化得到的json string: %s\n", string(tString))
	var t = &T{}
	if err := json.Unmarshal(tString, t); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("反序列化的gloang结构体: %#v\n", t)
}
