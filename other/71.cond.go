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
	cond.L.Lock()
	//等待通知,阻塞在此
	cond.Wait()
	fmt.Println(x)
	time.Sleep(time.Second * 1)
	defer func() {
		cond.L.Unlock() //释放锁
		waitgroup.Done()
	}()
}

func main() {
	for i := 0; i < 5; i++ {
		go test(i)
		waitgroup.Add(1)
	}
	fmt.Println("start all")
	time.Sleep(time.Second * 1)
	// 下发一个通知给已经获取锁的goroutine
	cond.Signal()
	time.Sleep(time.Second * 1)
	// 下发一个通知给已经获取锁的goroutine
	cond.Signal()
	time.Sleep(time.Second * 1)
	//下发广播给所有等待的goroutine
	fmt.Println("start Broadcast")
	cond.Broadcast()
	waitgroup.Wait()
}

// start all
// 1
// 0
// start Broadcast
// 4
// 2
// 3
