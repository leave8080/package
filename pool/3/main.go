package main

import (
	"log"
	"sync"
	"time"
)

//定义一个Worker接口,有个必须实现的Task()方法
type Worker interface {
	Task()
}

//定义一个类型Pool,有两个成员
type Pool struct {
	//成员work,通道类型,传递的是Worker类型
	work chan Worker
	//成员wg是计数信号量
	wg sync.WaitGroup
}

//定义New方法,返回的是Pool实例
//传递的参数是goroutine池的数量
func New(size int) *Pool {
	//实例化Pool类型
	pool := Pool{
		work: make(chan Worker),
	}
	//增加计数信号量
	pool.wg.Add(size)
	//使用循环创建多个goroutine
	for i := 0; i < size; i++ {
		//启动goroutine
		go func() {
			//从通道中获取值,这里如果没有会一直阻塞
			//这里会无限循环遍历,除非通道关闭了,否则不会跳出当前这个goroutine
			for w := range pool.work {
				//调用Worker类型的Task()方法
				w.Task()
			}
		}()
		pool.wg.Done()
	}
	return &pool
}

//给Pool类型定义Run方法
//参数是Worker类型
func (p *Pool) Run(w Worker) {
	//把Worker传进通道里
	p.work <- w
}

//给Pool类型定义 Shutdown方法
func (p *Pool) Shutdown() {
	//关闭通道
	close(p.work)
	//等待所有goroutine执行结束
	p.wg.Wait()
}

//定义一个字符串数组
var names = []string{
	"zhangsan",
	"lisi",
	"wangwu",
}

//定义一个类型namePrinter
type namePrinter struct {
	//成员name ,字符串类型
	name string
}

//给类型实现Worker接口
func (np *namePrinter) Task() {
	//打印namePrinter类型的name成员
	log.Printf(np.name)
	//睡眠一秒
	time.Sleep(time.Second)
}
func main() {
	//创建2个goroutine的池,因为通道是空的,这个地方有两个goroutine会阻塞在那
	pool := New(2)
	//定义计数信号量
	var wg sync.WaitGroup
	//增加计数,100次乘以数组元素个数
	wg.Add(100 * len(names))
	//循环100次,这个地方会瞬间生成300个goroutine,大并发的去执行任务
	for i := 0; i < 100; i++ {
		//循环数组
		for _, name := range names {
			//实例化namePrinter类型
			np := namePrinter{
				name: name,
			}
			//启动一个goroutine
			go func() {
				//调用Pool类型的run方法
				//传递的是Woker类型,因此要取地址
				//这里会把该Worker类型,发送到通道里,如果通道不为空,就会阻塞住
				//当300个goroutine,把name传递给run方法,会因为通道不为空被阻塞住
				//通道何时才能为空呢,也就只有在工作池里的goroutine把通道读走
				//因此会每次两个两个的打印,最多只会等待两个工作的完成
				pool.Run(&np)
				wg.Done()
			}()
		}
	}
	//等待上面的100次遍历结束
	wg.Wait()
	//停止工作池,关闭通道
	pool.Shutdown()
}
