package main

// 当使用通道作为函数的参数时，你可以指定这个通道是不是只用来发送或者接收值。
// 就是制定这个通道是只读还是只写
// 这个特性提升了程序的类型安全性

import "fmt"
// ping 函数定义了一个只允许发送数据的通道。
// 尝试使用这个通道来接收数据将会得到一个编译时错误。
// pings是只允许发送的
func ping(pings chan<- string, msg string) {
	pings <- msg
}
// pong 函数允许通道（pings）来接收数据，
// pings来接收通道的数据    通道的信息导入  <-chan 箭头在前  只写（从通道写入）
// pongs来发送通道的数据    通道的信息接收  chan<- 箭头在后  只读（读取到通道）
func pong(pings <-chan string, pongs chan<- string) {
	// channel通过操作符<-来接收和发送数据
	// 箭头在前面的就是接收数据
	// 箭头在后面的就是发送数据
	// ch <- v    // 箭头在后面就是发送数据   发送v到channel ch
	// v := <-ch  // 箭头在前面就是接收数据   从ch中接收数据，并赋值给v

	// pong 函数允许通道（pings）来接收数据，另一通道（pongs）来发送数据。
	// pings只写
	msg := <-pings
	// pongs只读
	pongs <- msg
}
func main() {
	// 新建立通道
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	// 
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// passed message
