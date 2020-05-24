package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	//全局变量
	counter int64
	//计数信号量
	wg sync.WaitGroup
	//mutex定义一段代码临界区
	mutex sync.Mutex
)

func main() {
	fmt.Println("hello")
	//计数加2,等待两个goroutine
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	//主goroutine等待子goroutine结束
	wg.Wait()
	fmt.Println("最终counter值:", counter)
}

//增加counter的值函数
func incCounter(i int) {
	//函数结束,减小信号量
	fmt.Println("go--->", i)
	defer wg.Done()
	for count := 0; count < 2; count++ {
		//创建这个临界区
		//同一时刻只允许一个goroutine进入
		mutex.Lock()
		//使用大括号只是为了让临界区看起来更清晰,并不是必须的
		{
			value := counter
			//强制调度器切换
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}
