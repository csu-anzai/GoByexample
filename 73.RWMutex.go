package main

import (
    "sync"
    "time"
)
// RWMutex读写锁
var m *sync.RWMutex

func main() {
    // 新建一个读写锁
    m = new(sync.RWMutex)
    
    // 多个同时读
    go read(1)
    go read(2)

    time.Sleep(2*time.Second)
}

func read(i int) {
    println(i,"read start")

    m.RLock()
    println(i,"reading")
    time.Sleep(1*time.Second)
    m.RUnlock()    

    println(i,"read over")
}
// PS D:\Github\Golang\Example> go run .\73.RWMutex.go
// 1 read start
// 1 reading
// 2 read start
// 2 reading
// 2 read over
// 1 read over