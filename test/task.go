package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个channel
	ch := make(chan string)

	//写数据协程
	go func() {
		fmt.Println("写协程开启:", time.Now())
		time.Sleep(2 * time.Second) //等待2秒
		ch <- "1"
		fmt.Println("写入成功:", time.Now())
	}()

	//读数据协程
	go func() {
		fmt.Println("读协程开启:", time.Now())
		i := <-ch
		fmt.Println("读取成功:", i, time.Now())
	}()

	//time.Sleep(5 * time.Second) //等待5秒,让所有协程跑完再退出
	fmt.Println("ok")
}

func Manger() {

}
