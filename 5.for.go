package main

import "fmt"

func main() {
	// :=表示的是不用定义类型
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		// 不要漏掉break啊 不然就死循环了
		break
	}
}
