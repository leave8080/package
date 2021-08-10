package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strconv"
)

type Redis struct {
	Redis *redis.Client
}

func InitRedis() (r *Redis) {
	var addr = "106.14.218.147:3309"
	var password = "mxcloud@2020"
	var db = 1

	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	p, err := c.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("failed to connect redis error:%+v", err))
	}
	fmt.Println("ping", p)
	r = &Redis{Redis: c}
	return r
}

type Info struct {
	FromRegion string `json:"from_region"`
	ToRegin    string `json:"to_regin"`
	ProductKey string `json:"product_key"`
	DeviceName string `json:"device_name"`
	WorkId     string `json:"work_id"`
}

func main() {
	rd := InitRedis()
	for i := 0; i < 100; i++ {
		info := &Info{
			FromRegion: "cn-shanghai",
			ToRegin:    "us",
			ProductKey: "ax123",
			DeviceName: "light",
			WorkId:     strconv.Itoa(i),
		}
		b, err := json.Marshal(info)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = rd.Redis.LPush("testList", b).Result()
		if err != nil {
			log.Println(err)
			return
		}
	}

}
