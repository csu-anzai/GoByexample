package main

import "fmt"

func main() {
	// 如果是堵塞channel，就会跑一个执行一个
	// 如果是buffchannel，就不会跟着那么紧
	jobs := make(chan int,5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	fmt.Println(<-done)
	// sent job 1
	// sent job 2
	// sent job 3
	// sent all jobs
	// received job 1
	// received job 2
	// received job 3
	// received all jobs
}
