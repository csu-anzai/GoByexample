package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// close之后的通道也是可以遍历到的，这个遍历就相当于顺序的同步的读取channel
	// 通道是没有index的
	// 如果没有close的话  会堵塞  死锁  因为读写缺少一个并且没有协程  所以死锁
	for elem := range queue {
		fmt.Println(elem)
	}
}

// one
// two
