package main

import (
	"gopro/patterns/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 25
	pooledResources = 2
)

//实现接口类型 资源类型
type dbConnection struct {
	ID int32
}

//实现接口方法
func (conn *dbConnection) Close() error {
	log.Printf("资源关闭,ID:%d\n", conn.ID)
	return nil
}

//给每个连接资源给id
var idCounter int32

//创建新资源
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Printf("创建新资源,id:%d\n", id)
	return &dbConnection{ID: id}, nil
}

//测试资源池
func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire(10 * time.Second)
	if err != nil {
		log.Println("获取资源超时")
		log.Println(err)
		return
	}
	//方法结束后将资源放进资源池
	defer p.Release(conn)
	//模拟使用资源
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("查询goroutine id:%d,资源ID：%d\n", query, conn.(*dbConnection).ID)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	//每个goroutine一个查询，每个查询从资源池中获取资源
	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	//主线程等待
	wg.Wait()
	log.Println("程序结束")
	//释放资源
	p.Close()
}
