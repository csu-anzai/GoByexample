package main

// 在这个例子中，我们将看到如何使用 Go 协程和通道实现一个工作池 。

import "fmt"
import "time"


func worker(id int, jobs <-chan int, results chan<- int) {
	// 因为job没有值  所以会等待
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second*3)
		results <- j * 2
	}
}
func main() {
	// 为了使用 worker 工作池并且收集他们的结果，我们需要2个通道。
	jobs := make(chan int,100)
	results := make(chan int,100)
	// worker是启动协程的，会控制协程的个数
	// jobs是任务
	// 会让几个worker进行执行jobs
	for w := 1; w <= 6; w++ {
		// w是协程个数
		// jobs是任务
		// resultes是执行堵塞的收集结果的
		go worker(w, jobs, results)
	}
	for j := 1; j <= 18; j++ {
		jobs <- j

	}
	close(jobs)


	// 堵塞的目的：让协程执行
	// 最后，我们收集所有这些任务的返回值。
	for a := 1; a <= 9; a++ {
		// fmt.Println("aaaaa")
		<-results
		// fmt.Println()
	}
	// close(results)

	// time.Sleep(time.Second*3)
}

// worker 10 processing job 1
// worker 2 processing job 3
// worker 3 processing job 4
// worker 7 processing job 5
// worker 6 processing job 6
// worker 5 processing job 2
// worker 4 processing job 7
// worker 8 processing job 8
// worker 9 processing job 9