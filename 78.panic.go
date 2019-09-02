package main

import (
	"fmt"
)
func main() {
	if err := test(); err != nil {
		fmt.Printf("mimic error: %v\n", err)	
	}
}
func mimicError(key string) error {
	return fmt.Errorf("mimic error: %s", key)	
}
func test() (err error) {
	defer func() {
		fmt.Println("start panic defer")
		// 如果recover方法被调用，但是没有任何的panic发生，recover方法只会返回nil	
		if r := recover(); r != nil {
			fmt.Println("defer panic:", r)
		}
	}()
	fmt.Println("start test")
	err = mimicError("1")
	fmt.Println("end test")
	
	return err
}
// start test
// end test
// start panic defer
// mimic error: mimic error: 1