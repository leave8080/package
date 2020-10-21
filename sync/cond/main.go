package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func test(x int) {
	cond.L.Lock()         //获取锁
	defer cond.L.Unlock() //释放锁
	fmt.Println("aaa: ", x)
	cond.Wait() //等待通知  暂时阻塞
	fmt.Println("bbb: ", x)
	time.Sleep(time.Second * 2)

}

func main() {
	for i := 0; i < 5; i++ {
		go test(i)
	}
	fmt.Println("start all")
	time.Sleep(time.Second * 1)
	fmt.Println("broadcast")
	cond.Signal() // 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	cond.Signal() // 3秒之后 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	cond.Broadcast() //3秒之后 下发广播给所有等待的goroutine
	time.Sleep(time.Second * 10)
	fmt.Println("finish all")

}

/*
众所周知，go语言在多线程方面的支持是十分完备的。在go语言sync包中提供了一个Cond类，这个类用于goroutine之间进行协作。

这个类并不复杂，只有三个函数，Broadcast() , Signal(), Wait()， 一个成员变量，L　Lock

其中Broadcast()实现的功能是唤醒在这个cond上等待的所有的goroutine，而Signal()则只选择一个进行唤醒。Wait()自然是让goroutine在

这个cond上进行等待了。这几个函数有以下几个注意点：

1.Wait()函数在调用时一定要确保已经获取了其成员变量锁L ,因为Wait第一件事就是解锁。　但是需要注意的是，当Ｗａｉｔ（）结束等待返回之前，

　它会重新对Ｌ进行加锁，也就是说，当Ｗａｉｔ结束，调用它的Ｇｏｒｏｕｔｉｎｅ仍然会获取Ｌｏｃｋ　Ｌ。

２．调用Ｂｒｏａｄｃａｓｔ（）函数会导致系统切换到之前在等待的那个Ｇｏｒｏｕｔｉｎｅ进行执行。


*/
