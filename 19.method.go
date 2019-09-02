package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	fmt.Println("area中r此时的状态", r) //area中r此时的状态 &{10 5}
	return r.width * r.height
}
func (r rect) perim() int {
	fmt.Println("perim中r此时的状态", r) //perim中r此时的状态 {10 5}
	return 2*r.width + 2*r.height
}
func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("main中r此时的状态", r) //&{10 5}
	// r此时是地址 地址对应的width变量*地址对应的height变量
	fmt.Println("带*的类型area: ", r.area())
	// r此时是具体值 直接进行计算
	fmt.Println("不带*的类型perim:", r.perim())

	rp := &r
	fmt.Println("rp此时的状态", rp)
	// rp此时是地址 area()的r为地址  直接带入即可
	fmt.Println("带*的类型area: ", rp.area())
	// rp是地址  perim()的r为值  需要进行转换  找出地址对应的变量值
	fmt.Println("不带*的类型perim:", rp.perim())
}
