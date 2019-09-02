package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}
func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	// 这个部分体现_的作用
	_, c := vals()
	fmt.Println(c)
}

// 3
// 7
// 7
