package main

import (
	"encoding/json"
	"fmt"
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
	var animals []Animal
	// 会按照animala的方式进行编码
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}

// [{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
