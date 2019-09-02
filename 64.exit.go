package main

// 使用 os.Exit 来立即进行带给定状态的退出。

import "fmt"
import "os"

func main() {
	// 当使用 os.Exit 时 defer 将不会 执行，
	// 所以这里的 fmt.Println将永远不会被调用。
	defer fmt.Println("!")
	// 退出并且退出状态为 3。
	os.Exit(3)
}

// // 如果你使用 go run 来运行 exit.go，那么退出状态将会被 go捕获并打印。
// $ go run exit.go
// exit status 3
// //使用编译并执行一个二进制文件的方式，你可以在终端中查看退出状态。
// $ go build exit.go
// $ ./exit
// $ echo $?
// 3
