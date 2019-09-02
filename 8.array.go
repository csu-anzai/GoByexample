package main

import "fmt"

func main() {
	// 一维数组的原始定义
	//
	// var 数组名 [个数]类型
	// 数组名:=[个数]类型{数值}
	//
	// go真的有些反常规啊  
	// 别的都是int a[5]={};
	//   而go是var a [5]int
	var a [5]int
	fmt.Println(a)

	// 一维数组的变形定义
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 定义一个二维数组
	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("二维数组：", c)
}
