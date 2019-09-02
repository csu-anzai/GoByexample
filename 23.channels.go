package main

import "fmt"

func main() {
	messages := make(chan string)

	// channel与goroutine必须并存
	// 如果只有channel，根本无法运行，因为全部死锁了
	// 如果只有goroutine，协程根本不会执行，因为没时间，主程序已经结束了
	// 因为有了堵塞，协程才会执行
	go func() {
		fmt.Println("2")
		messages <- "aaaa"
	}()

	// channel等待写入数据，此处堵塞
	// 协程写入数据后，会继续执行
	fmt.Println("1")
	msg := <-messages

	fmt.Println("3")

	fmt.Println(msg)
}

// ping
