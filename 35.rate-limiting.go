package main

import "time"
import "fmt"

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)
	for req := range requests {
		// 每200ms执行一次
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()
	burstyRequests := make(chan int, 6)
	for i := 1; i <= 6; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		// 堵塞了200ms
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
// request 1 2019-09-03 17:57:00.0845963 +0800 CST m=+0.211446401
// request 2 2019-09-03 17:57:00.2842718 +0800 CST m=+0.411121901
// request 3 2019-09-03 17:57:00.4847832 +0800 CST m=+0.611633301
// request 4 2019-09-03 17:57:00.6843482 +0800 CST m=+0.811198301
// request 5 2019-09-03 17:57:00.8858815 +0800 CST m=+1.012731601

// request 1 2019-09-03 17:57:00.8858815 +0800 CST m=+1.012731601
// request 2 2019-09-03 17:57:00.8864562 +0800 CST m=+1.013306401
// request 3 2019-09-03 17:57:00.8864562 +0800 CST m=+1.013306401
// request 4 2019-09-03 17:57:01.0859668 +0800 CST m=+1.212816901
// request 5 2019-09-03 17:57:01.2863722 +0800 CST m=+1.413222401
// request 6 2019-09-03 17:57:01.4860956 +0800 CST m=+1.612945701