package main

import "fmt"

// 此处的赋值运用了  值传递
func zeroval(ival int) {
	ival = 0
}

// 这里的*int表示的是变量 不是地址
func zeroptr(iptr *int) {
	//iptr本身是一个地址  *iptr就变成了变量
	*iptr = 0
	fmt.Println("pointer:",iptr)
}
func main() {
	i := 1
	//
	// 值传递
	//
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	//
	// 引用传递
	//
	// zeroptr 有一和上面不同的 *int 参数，意味着它用了一
	// 个 int指针。函数体内的 *iptr 接着解引用 这个指针，从它
	// 内存地址得到这个地址对应的当前值。对一个解引用的指针赋值
	// 将会改变这个指针引用的真实地址的值。
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	// 返回的地址
	fmt.Println("pointer:", &i)
}

// initial: 1
// zeroval: 1
// pointer: 0xc042052058
// zeroptr: 0
// pointer: 0xc042052058
