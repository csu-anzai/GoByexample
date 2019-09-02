// Go 支持在结构体类型中定义方法 。
package main

import "fmt"

// 接收器类型rect
type rect struct {
	width, height int
}

// 这里的 area 方法有一个接收器类型 rect。此处的接收器为r
// 可以为值类型或者指针类型的接收器定义方法。这里是一个值类型接收器的例子。
// func funcName(input1 type1, input2 type2) (otput1 type1, output2 type2) {
// 		这里是处理逻辑代码
// 		返回多个值
// 		return value1, value2
// }
// 由此可以看出在普通的函数前面添加了一个接收器对象r
// 这里是为接收器r的方法area 返回的是int的类型
func (r *rect) area() int {
	return r.width * r.height
}

// 这里是为接收器r的方法perim 返回的是int的类型
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	// r是一个对象
	// 这个对象在不同的方法中 会有不同的结果
	r := rect{width: 10, height: 5}
	// 这里我们调用上面为结构体定义的两个方法。
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go 自动处理方法调用时的值和指针之间的转化。
	// 你可以使用指针来调用方法来避免在方法调用时产生一个拷贝，或者让方法能够改变接受的数据。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

// area:  50
// perim: 30
// area:  50
// perim: 30
