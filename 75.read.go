package main

import (
    "sync"
    "time"
)

var m *sync.RWMutex

func main() {
    m = new(sync.RWMutex)
    go write(1)
    go read(3)
    time.Sleep(20 * time.Second)
}

func read(i int) {
    println(i, "read start")
    m.RLock()
    var p = 0
    var pr = "read"
    for {
        pr += "."
        if p == 3 {
            break
        }
        time.Sleep(350 * time.Millisecond)
        p++
        println(i, pr)

    }
    m.RUnlock()
    println(i, "read end")
}

func write(i int) {
    println(i, "write start")

    m.Lock()
    var p = 0
    var pr = "write"
    for {
        pr += "."
        if p == 3 {
            break
        }
        time.Sleep(350 * time.Millisecond)
        p++
        println(i, pr)

    }
    m.Unlock()
    println(i, "write end")
}
// 1 write start
// 3 read start
// 1 write.
// 1 write..
// 1 write...
// 1 write end
// 3 read.
// 3 read..
// 3 read...
// 3 read end