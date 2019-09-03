package main

import (
    "fmt"
    "sync"
    "time"
)

func ccc(i int)(){
		mutex.Lock()
		fmt.Println("Lock:", i)
		time.Sleep(time.Second)
		fmt.Println("Unlock:", i)
		mutex.Unlock()
		defer wait.Done()
}
var mutex sync.Mutex
var wait = sync.WaitGroup{}

func main() {

    fmt.Println("互斥锁")
    mutex.Lock()

    for i := 1; i <= 3; i++ {
        wait.Add(1)

        go ccc(i)
    }
    mutex.Unlock()

    wait.Wait()
}


// 互斥锁锁定一个资源
// 其他的协程无法进行操作此段的资源  其他的协程都选择等待此段锁代码的结束
// 这个资源已经被当前的协程进行锁定，其他的协程也要操作此资源的都要等待
// 锁是表明一段代码中只能同步执行的
// 互斥锁
// Lock: 1
// Unlock: 1
// Lock: 3
// Unlock: 3
// Lock: 2
// Unlock: 2