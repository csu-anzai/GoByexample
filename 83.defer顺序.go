package main

import (
	"fmt"
)

func main(){
	// defer是栈的结构  后入先出
	defer fmt.Println("1")
	defer fmt.Println("2")
}