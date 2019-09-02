package main

import "fmt"

// Go 协程 在执行上来说是轻量级的线程。

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 普通的情况输出
	// 假设我们有一个函数叫做 f(s)。我们使用一般的方式调并同时运行。
	f("direct")

	// 协程的情况输出
	// 使用 go f(s) 在一个 Go 协程中调用这个函数。这个新的 Go 协程将会并行的执行这个函数调用。
	go f("goroutine")
	// 你也可以为匿名函数启动一个 Go 协程。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//
	// 现在这两个 Go 协程在独立的 Go 协程中异步的运行，所以我们需要等它们执行结束。
	// 这里的 Scanln 代码需要我们在程序退出前按下任意键结束。
	//
	// 当我们运行这个程序时，
	// 将首先看到阻塞式调用的输出，然后是两个 Go 协程的交替输出。
	// 这种交替的情况表示 Go 运行时是以异步的方式运行协程的。
	//
	// 总结：就是执行没有结束前 按下任意的键 协程就会结束
	// 如果不按的话  协程会并行执行
	// 这里面的并行指的是go f()函数与go func()函数的执行
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

// 一结果般方式的输出
// direct : 0
// direct : 1
// direct : 2
//
// goroutines的并行方式
// goroutine : 0
// going
// goroutine : 1
// goroutine : 2
// <enter>
// aaaaa
// done
