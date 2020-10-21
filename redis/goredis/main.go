package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
	Desc string `json:"desc"`
}

func main() {
	c, err := redis.Dial("tcp", "192.168.1.100:6379")
	if err != nil {
		fmt.Println("conn err")
		return
	}
	defer c.Close()

	//_, err = c.Do("set", "Q1", "3333", "EX", 5)
	//if err != nil {
	//	fmt.Println("set err")
	//	return
	//}
	//r, err := redis.String(c.Do("get", "Q1"))
	//if err != nil {
	//	fmt.Println("get err")
	//	return
	//}

	//_, err = c.Do("hset", "Q2", "3333", "tcp", "127.0.0.0.1", "EX", 5)
	//if err != nil {
	//	fmt.Println("set err", err.Error())
	//	return
	//}
	//r, err := redis.String(c.Do("hget", "Q2"))
	//if err != nil {
	//	fmt.Println("get err")
	//	return
	//}
	//
	//fmt.Println(r)
	//cacheName := "*:20180718"
	//keys, err := redis.Strings(c.Do("keys", cacheName))
	//
	//for _, key := range keys {
	//	activeInfo, err := redis.IntMap(c.Do("hgetall", key))
	//	if err != nil {
	//		fmt.Errorf("查询单个key的日活信息出错：err:%v\n", err)
	//		return
	//	}
	//	fmt.Println("key-%v, activeInfo-%+v\n", key, activeInfo)
	//}
	stu1 := Student{
		Name: "st1",
		Age:  20,
		Sex:  1,
		Desc: "dasdad",
	}
	var ub []byte
	ub, err = json.Marshal(stu1)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	c.Send("set", "st1", ub)
	c.Flush()

	v, err := c.Receive()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("%#v\n", v)

	c.Send("GET", "st1")
	c.Flush()
	var vb []byte
	vb, err = redis.Bytes(c.Receive())
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	//fmt.Printf("%#v\n", vb)
	u := new(Student)
	// 直接将　user set,再Unmarshal为结构体时会失败
	err = json.Unmarshal(vb, &u)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("%#v\n", u)

}
