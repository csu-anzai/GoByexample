package main

import "fmt"

// channel是协程与协程之前的通道
//
// 在并发的 Go 程序中的 Go 协程的辅助特性：通道。
// 通道 是连接多个 Go 协程的管道。你可以从一个 Go 协程将值发送到通道，然后在别的 Go 协程中接收。

func main() {
	//
	// 通道的类型就是传递值的类型
	//
	// 新建一个通道
	// 使用 make(chan val-type) 创建一个新的通道。通道类型就是他们需要传递值的类型。
	// 此时 messages是一个通道
	messages := make(chan string)
	// 赋初值
	// 使用 channel <- 语法 发送 一个新的值到通道中。这里我们在一个新的 Go 协程中
	// 发送 "ping" 到上面创建的messages 通道中。
	go func() { 
		messages <- "ping" 
	}()
	// 使用 <-channel 语法从通道中 接收 一个值。这里将接收我们在上面发送的 "ping" 消息并打印出来。
	// 通道的初值赋值给一个变量
	//
	// 通道的传值使用<-
	//
	msg := <-messages
	fmt.Println(msg)
}

// ping
