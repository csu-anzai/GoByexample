package main

import (
	"encoding/json"
	"fmt"
)

type Change struct {
	Mid     int      //菜单Id
	Actions []string //拥有的权限 "add"  "view"  "delete"  "update"
}
type Change_slice struct {
	ChgArr []Change //一个角色对应的菜单以及权限
}

func main() {
	// 序列化
	//对象序列化为json字符串---------------------------------Begin
	var c1, c2 Change
	var msg Change_slice
	c1.Mid = 1
	c1.Actions = []string{"view", "add"}
	c2.Mid = 2
	c2.Actions = []string{"delete", "add", "update"}
	msg.ChgArr = []Change{c1, c2}
	fmt.Println(msg)
	b, err := json.Marshal(msg)
	if err == nil {
		fmt.Println(string(b))
		fmt.Println("1")
	}
	//对象序列化为json字符串---------------------------------End

	//json字符串反序列化为对象---------------------------------Begin
	var str string = `{"ChgArr":[{"Mid":1,"Actions":["view","add"]},{"Mid":2,"Actions":["delete","add","update"]}]}`
	msgs := str
	// var msgs Change_slice
	e := json.Unmarshal([]byte(str), &msgs)
	if err != nil {
		fmt.Println("Can't decode json message", err)
	} else {
		fmt.Println("\n")
		fmt.Println(e)
		fmt.Println("aaaaa",str)
		fmt.Println(2)
	}
	//json字符串反序列化为对象---------------------------------End
}

// {}就是结构体
// []就是切片
// {[{1 [view add]} {2 [delete add update]}]}
// {"ChgArr":[{"Mid":1,"Actions":["view","add"]},{"Mid":2,"Actions":["delete","add","update"]}]}
// 1


// json: cannot unmarshal object into Go value of type string
// {"ChgArr":[{"Mid":1,"Actions":["view","add"]},{"Mid":2,"Actions":["delete","add","update"]}]}
// 2