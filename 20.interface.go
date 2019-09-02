package main

import "fmt"

// import "math"

// 接口：一个事物多个方法
// 抽象：多个事物一个特性

// 在我们的例子中，我们将让 rect 和 circle 实现这个接口
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// 这里是一个几何体的基本接口。
// 接口全部是方法
// 结构体是数据类型的集合
// 接口是method的集合 再细分的话就是  有对象的func的集合
type geometry interface {
	// 面积
	area() float64
	// 周长
	perim() float64
}

// 要在 Go 中实现一个接口，我们只需要实现接口中的所有方法。这里我们让 rect 实现了 geometry 接口。
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// 如果一个变量的是接口类型，那么我们可以调用这个被命名的接口中的方法。
// 这里有一个一通用的 measure 函数，利用这个特性，它可以用在任何 geometry 上。
// 集合了两种方法的方法
//
// 接口也是一种类型geometry 对应的对象就是g
//
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {
	r := rect{width: 3, height: 4}
	// 结构体类型 circle 和 rect 都实现了 geometry接口，所以我们可以使用它们的实例作为 measure 的参数。
	measure(r)
}

// {3 4}
// 12
// 14
// {5}
// 78.53981633974483
// 31.41592653589793
