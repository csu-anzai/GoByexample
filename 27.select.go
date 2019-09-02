package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		fmt.Println(1)
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		fmt.Println(5)
		time.Sleep(time.Second * 5)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

// 我们首先接收到值 "one"，然后就是预料中的 "two"了。
// 注意从第一次和第二次 Sleeps 并发执行，总共仅运行了两秒左右。
// received one
// received two
