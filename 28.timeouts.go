package main

import "time"
import "fmt"

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	for i:=0;i<2;i++{
		select {
		case res := <-c1:
			fmt.Println(res)
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1")
		}
	}

}

// timeout 1
// result 1
