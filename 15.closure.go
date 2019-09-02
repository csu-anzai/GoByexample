package main

import "fmt"

// intSeq表示的函数的名称
// func()表示的返回的匿名函数 括号里面应该是输入的数据的类型
// int是返回的匿名函数的类型  这个返回的函数使用闭包的方式 隐藏 变量 i。
//
// 返回的数据类型func() int 其实里面是一个匿名函数  函数没有名字
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func main() {
	// 这个函数的值包含了自己的值 i，这样在每次调用 nextInt 时都会更新 i 的值。
	nextInt := intSeq()
	// 1
	fmt.Println(nextInt())
	// 2
	fmt.Println(nextInt())
	// 3
	fmt.Println(nextInt())
	newInts := intSeq()
	// 1
	fmt.Println(newInts())
}

// 1
// 2
// 3
// 1
