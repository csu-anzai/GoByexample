package main

import "fmt"

// 后面的int漏了
// 这里在表明func返回值的类型
//
// func 名称(a 类型,b 类型)返回类型{}
//
func plus(a int, b int) int {

	return a + b
}

func main() {
	res := plus(1, 2)
	fmt.Println(res)
}
