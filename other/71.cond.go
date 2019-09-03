package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)
var waitgroup sync.WaitGroup

func test(x int) {
	//获取锁
	cond.L.Lock()
	//等待通知,阻塞在此
	// func (c *Cond) Wait()：挂起goroutine的操作
	cond.Wait()
	fmt.Println(x)
	time.Sleep(time.Second * 1)
	defer func() {
		cond.L.Unlock() 	//释放锁
		waitgroup.Done()	//减一个协程
	}()
}
// cond其实就是哪个控制协程执行的包函数
// Signal是执行一个协程的信号
// Broadcast是执行全部协程的信号
// wait 协程等待通知,阻塞在此
// NewCond创建条件
// cond.L.Lock() 给协程加锁
// cond.L.Unlock() 	//释放锁

func main() {
	// func NewCond(l Locker) *Cond：
	// 用于创建条件，根据实际情况传入sync.Mutex或者sync.RWMutex的指针，一定要是指针，否则会发生复制导致锁的失效
	// 开5个协程
	for i := 0; i < 5; i++ {
		go test(i)
		waitgroup.Add(1)
	}
	fmt.Println("start all")
	time.Sleep(time.Second * 1)
	// 下发一个通知给已经获取锁的goroutine
	// func (c *Cond) Signal()：随机唤醒等待队列上的goroutine，随机的方式效率更高
	cond.Signal()
	time.Sleep(time.Second * 1)

	cond.Signal()
	time.Sleep(time.Second * 1)
	//下发广播给所有等待的goroutine
	// func (c *Cond) Broadcast()：唤醒条件上的所有goroutine
	fmt.Println("start Broadcast")
	cond.Broadcast()
	// 协程数为0就会完成
	waitgroup.Wait()
}

// start all
// 1
// 0
// start Broadcast
// 2
// 3
// 4