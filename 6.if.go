package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7是偶数")
	} else {
		fmt.Println("7是奇数")
	}
	// 可以进行赋值
	if num := 9; num%2 == 1 {
		fmt.Println("num是奇数")
	} else {
		fmt.Println("num是偶数")
	}
}
