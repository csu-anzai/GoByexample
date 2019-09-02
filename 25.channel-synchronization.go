package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("2")
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
func main() {
	done := make(chan bool, 1)
	go worker(done)

	fmt.Print("1")
	// 等待吸入channel数据
	fmt.Println(<-done)
}

// 12working...done
// true

