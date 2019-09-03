package main
 
import (
	"log"
	"sync"
)
 
// pool就是一个中间件
// put放入
// get取出
// New为默认

func main(){
	// 建立对象
	// var pipe = sync.Pool{
	// 	New:func()interface{}{
	// 		var a string
	// 		a="aaaaaa"
	// 		return &a
	// 	},
	// }

	var pipe = &sync.Pool{
		New: func() interface{} {
			var mapp map[string]float32;
			mapp = make(map[string]float32);
			mapp["name"]=11.11;

			return mapp
		},
	}


	// 准备放入的字符串
	val := "Hello,World!"
	// 放入
	pipe.Put(val)
	// 取出
	log.Println(pipe.Get())
	// 再取就没有了,会自动调用NEW
	log.Println(pipe.Get())
}
//  有new的时候
// 2018/12/29 17:50:10 Hello,World!
// 2018/12/29 17:50:11 Hello,BeiJing

// 没有new的时候
// 2019/09/03 19:40:11 Hello,World!
// 2019/09/03 19:40:11 <nil>