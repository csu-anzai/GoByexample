package main

// 互斥锁（英语：英语：Mutual exclusion，缩写 Mutex）
// 是一种用于多线程编程中，
// 防止两条线程同时对同一公共资源（比如全局变量）
// 进行读写的机制
import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// 在前面的例子中，我们看到了如何使用原子操作来管理简单的计数器。
// 对于更加复杂的情况，我们可以使用一个
// 互斥锁来在 Go 协程间安全的访问数据。

func main() {
	// 在我们的例子中，state 是一个 map。
	var state = make(map[int]int)
	// 这里的 mutex 将同步对 state 的访问。
	var mutex = &sync.Mutex{}
	// we'll see later, ops will count how many operations we perform against the state.
	// 为了比较基于互斥锁的处理方式和我们后面将要看到的其他方式，
	// ops 将记录我们对 state 的操作次数。
	var ops int64 = 0
	// 这里我们运行 100 个 Go 协程来重复读取 state。
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				// 每次循环读取，我们使用一个键来进行访问，Lock() 这个 mutex 来确保对 state 的独占访问，读取选定的键的值，Unlock() 这个mutex，并且 ops 值加 1。
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				// 为了确保这个 Go 协程不会在调度中饿死，我们在每次操作后明确的使用 runtime.Gosched()进行释放。这个释放一般是自动处理的，像例如每个通道操作后或者 time.Sleep 的阻塞调用后相似，但是在这个例子中我们需要手动的处理。
				runtime.Gosched()
			}
		}()
	}
	// 同样的，我们运行 10 个 Go 协程来模拟写入操作，使用和读取相同的模式。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	// 让这 10 个 Go 协程对 state 和 mutex 的操作运行 1 s。
	time.Sleep(time.Second)
	// 获取并输出最终的操作计数。
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
	// 对 state 使用一个最终的锁，显示它是如何结束的。
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}

// [ `go run 37.mutexes.go` | done: 1.939s ]
// 	ops: 3911414
// 	state: map[0:21 4:37 3:62 2:77 1:54]
