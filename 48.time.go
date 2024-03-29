package main

import "fmt"
import "time"

// Go 对时间和时间段提供了大量的支持；这里是一些例子。
//
// 2012-10-31 15:50:13.793654 +0000 UTC
// 2009-11-17 20:34:58.651387237 +0000 UTC
// 2009
// November
// 17
// 20
// 34
// 58
// 651387237
// UTC
// Tuesday
// true
// false
// false
// 25891h15m15.142266763s
// 25891.25420618521
// 1.5534752523711128e+06
// 9.320851514226677e+07
// 93208515142266763
// 2012-10-31 15:50:13.793654 +0000 UTC
// 2006-12-05 01:19:43.509120474 +0000 UTC

func main() {
	p := fmt.Println
	// 得到当前时间。
	now := time.Now()
	p(now)
	// 通过提供年月日等信息，你可以构建一个 time。时间总是关联着位置信息，例如时区。
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	// 你可以提取出时间的各个组成部分。
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	// 输出是星期一到日的 Weekday 也是支持的。
	p(then.Weekday())
	// 这些方法来比较两个时间，分别测试一下是否是之前，之后或者是同一时刻，精确到秒。
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))
	// 方法 Sub 返回一个 Duration 来表示两个时间点的间隔时间。
	diff := now.Sub(then)
	p(diff)
	// 我们计算出不同单位下的时间长度值。
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	// 你可以使用 Add 将时间后移一个时间间隔，或者使用一个 - 来将时间前移一个时间间隔。
	p(then.Add(diff))
	p(then.Add(-diff))
}
