package main

import "fmt"

type person struct {
	name string
	age  int
}

var a int
func main() {
	a=1

	fmt.Println(&a)
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	// 下面的输出结果是	&{Ann 40}
	fmt.Println(&person{name: "Ann", age: 40})
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	// & 前缀生成一个结构体指针。
	// 使用点来访问结构体字段。
	// 也可以对结构体指针使用. - 指针会被自动解引用。
	// 结构体是可变的。

	// 进行传递的地址
	sp := &s
	fmt.Println("sp",sp)
	fmt.Println(sp.age)
	// 下面是在进行修改sg.age的值
	sp.age = 51
	fmt.Println(sp.age)
}

// {Bob 20}
// {Alice 30}
// {Fred 0}
// &{Ann 40}
// Sean
// 50
// 51
