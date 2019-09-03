package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	m = new(sync.RWMutex)

	// 写的时候啥也不能干
	go write(1)
	go read(2)
	go write(3)

	time.Sleep(7 * time.Second)
}

// RWMUX是控制协程对于资源的改变顺序的
// RLOCK RULOCK
// LOCK ULOCK
// 在一个协程里面
// 对资源的
// 读锁的时候别的也可以读，但是不可以写
// 写锁的时候都不可以操作，不可以读也不可以写

func read(i int) {
	println(i, "read start")

	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()

	println(i, "read over")
}

func write(i int) {
	time.Sleep(1 * time.Second)

	println(i, "write start")

	m.Lock()
	println(i, "writing")
	time.Sleep(2*time.Second)
	m.Unlock()

	println(i, "write over")
}

// 2 read start
// 2 reading
// 3 write start
// 1 write start
// 2 read over
// 3 writing
// 3 write over
// 1 writing
// 1 write over