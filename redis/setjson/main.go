package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var redisConc *RedisConnc

// User ...
type User struct {
	Name string `json:"user"`
	Age  int    `json:"age"`
}

// RedisConnc ...
type RedisConnc struct {
	pool *redis.Pool
}

func main() {
	var err error
	InitRedis("192.168.11.34:6379", "123456")
	c := redisConc.pool.Get()
	user := User{
		Name: "test-name",
		Age:  18,
	}
	defer c.Close()
	//　1.直接将　user set,再Unmarshal为结构体时会失败
	// c.Send("SET", "user", user)
	// 2. 应该先 使用Marshal序列化
	var ub []byte
	ub, err = json.Marshal(user)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	c.Send("SET", "user", ub)
	c.Flush()
	v, err := c.Receive()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("%#v\n", v)

	c.Send("GET", "user")
	c.Flush()
	var vb []byte
	vb, err = redis.Bytes(c.Receive())
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("%#v\n", vb)
	u := new(User)
	// 直接将　user set,再Unmarshal为结构体时会失败
	err = json.Unmarshal(vb, &u)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("%#v\n", u)
}

// InitRedis ...
func InitRedis(host string, passwd string) error {
	var pool = &redis.Pool{
		MaxIdle:     50,
		MaxActive:   100,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {

			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			if passwd != "" {
				if _, err := c.Do("AUTH", passwd); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		// custom connection test method
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if _, err := c.Do("PING"); err != nil {
				return err
			}
			return nil
		},
	}
	redisConc = &RedisConnc{pool}
	return nil
}

/*
Pipelining(管道)
管道操作可以理解为并发操作，并通过Send()，Flush()，Receive()三个方法实现。客户端可以使用send()方法一次性向服务器发送一个或多个命令，命令发送完毕时，使用flush()方法将缓冲区的命令输入一次性发送到服务器，客户端再使用Receive()方法依次按照先进先出的顺序读取所有命令操作结果。

Send(commandName string, args ...interface{}) error
Flush() error
Receive() (reply interface{}, err error)
Send：发送命令至缓冲区
Flush：清空缓冲区，将命令一次性发送至服务器
Recevie：依次读取服务器响应结果，当读取的命令未响应时，该操作会阻塞。
*/
