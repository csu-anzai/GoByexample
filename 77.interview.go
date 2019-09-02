package main

// 在这个例子中，我们将看到如何使用 Go 协程和通道实现一个工作池 。

import "fmt"



func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()
 
		for elem := range set.s {
			ch <- elem
		}
 
		close(ch)
		set.RUnlock()
 
	}()
	return ch
}