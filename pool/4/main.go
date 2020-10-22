package main

import (
	"fmt"
	"sync"
	"time"
)

type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

// 创建并发控制池, 设置并发数量与任务总数量
func NewPool(cap, total int) *Pool {
	if cap < 1 {
		cap = 1
	}
	p := &Pool{
		queue: make(chan int, cap),
		wg:    new(sync.WaitGroup),
	}
	p.wg.Add(total)
	return p
}

// 向并发队列中添加一个
func (p *Pool) AddOne() {
	p.queue <- 1
}

// 并发队列中释放一个, 并从任务总数量中减去一个
func (p *Pool) DelOne() {
	<-p.queue
	p.wg.Done()
}

func main() {
	//开启多核处理
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//需要处理的下载任务
	urls := []string{
		"01", "02", "03", "04", "05", "06",
		"07", "08", "09", "10", "11", "12",
		"13", "14", "15", "16", "17", "18",
		"19", "20", "21", "22", "23", "24",
		"25", "26", "27", "28", "29", "30",
	}
	//初始化一个控制池,设置并发数量20
	pool := NewPool(20, len(urls))
	//计算执行时间
	now := time.Now()
	//并发处理
	for _, v := range urls {
		go func(url string) {
			pool.AddOne() // 向并发控制池中添加一个, 一旦池满则此处阻塞
			//任务处理
			err := Download(url)
			if nil != err {
				println(err)
			}
			pool.DelOne() // 从并发控制池中释放一个, 之后其他被阻塞的可以进入池中
		}(v)
	}
	//等待所有下载全部完成
	pool.wg.Wait()
	//计算执行时间
	next := time.Now()
	fmt.Println("执行时间:", next.Sub(now))
}

func Download(s string) error {
	//下载测试
	time.Sleep(2 * time.Second)
	//这里下载可以做一些并发时锁的使用
	println("任务ID:", s)
	return nil
}
