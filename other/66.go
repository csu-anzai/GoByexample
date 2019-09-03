package main

import (
	"fmt"
	"time"
)

func produce(p chan<- int) {
	for i := 0; i < 3; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}
func consumer(c <-chan int) {
	for i := 0; i < 3; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}
func main() {
	ch := make(chan int, 10)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}

// send: 0
// send: 1
// send: 2
// receive: 0
// receive: 1
// receive: 2
