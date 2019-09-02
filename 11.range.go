package main

import "fmt"

func main() {

	// for range用来返回切片
	nums := []int{2, 3, 4}
	sum := 0
	// 这里的下划线表示：键值不需要返回
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// for range用来返回map
	// map[string]string{}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

// sum: 9
// index: 1
// a -> apple
// b -> banana
// 0 103
// 1 111
// g是71  o是79
