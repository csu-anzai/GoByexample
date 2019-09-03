package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Animal struct {
	Name  string
	Order string
	site  string
}

func main() {

	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)
	a:=reflect.TypeOf(jsonBlob)
	fmt.Println(a)
	// 输出来的全部是ascii码
	fmt.Println(jsonBlob)

	var animals []Animal
	// 会按照animala的方式进行编码
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	// %+v会打印出键来
	fmt.Printf("%+v", animals)
}

// [{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
