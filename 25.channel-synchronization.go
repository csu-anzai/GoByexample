package main
// 通道同步
// 我们可以使用通道来同步 Go 协程间的执行状态。
// 这里是一个使用阻塞的接受方式来等待一个 Go 协程的运行结束。
import "fmt"
import "time"
// 这是一个我们将要在 Go 协程中运行的函数。
// done 通道将被用于通知其他 Go 协程这个函数已经工作完毕。
//
// 样式的定义：func sum(c chan int) {}
// done是一个channel 类型是int
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// 发送一个值来通知我们已经完工啦。
	// 通道进行赋值
	done <- true
}
func main() {
	// 运行一个 worker Go协程，并给予用于通知的通道。
	done := make(chan bool, 1)
	go worker(done)
	// 程序将在接收到通道中 worker 发出的通知前一直阻塞。
	// 此处的done通道就是在表示没结束前 一直堵塞  不执行其他的

	// 此处done是没有赋值的  所以一直是堵塞的状态
	// 当go worker准备好done的数值之后进行执行输出
	// <-done
	fmt.Println(<-done)
}

// working...done
// true

// 如果你把 <- done 这行代码从程序中移除，程序甚至会在 worker还没开始运行时就结束了。
