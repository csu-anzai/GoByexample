package main

import (
    "fmt"
    "time"
)
func produce(p chan<- int) {
    for i := 0; i < 4; i++ {
        p <- i
        fmt.Println("send:", i)
    }
}
func consumer(c <-chan int) {
    for i := 0; i < 4; i++ {
        v := <-c
        fmt.Println("receive:", v)
    }
}
func main() {
    ch := make(chan int)
    go produce(ch)
    go consumer(ch)
    time.Sleep(1 * time.Second)
}
// PS D:\Github\Golang\Example> go run .\65.go
// receive: 0
// send: 0
// send: 1
// receive: 1
// receive: 2
// send: 2
// send: 3
// receive: 3