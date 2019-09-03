// 用once可以保证上面的oncebody被执行一次
package main
 
import (
	"fmt"
	"sync"
	"time"
)
 
func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 5; i++ {
		j := i
		go func(int) {
			once.Do(onceBody)
			fmt.Println(j)
			done <- true
		}(j)
	}
	fmt.Println(<-done)
	// 表示等待所有的进行执行完成
	time.Sleep(2 * time.Second)
}
 
// Only once
// 0
// 4
// 1
// 2
// 3