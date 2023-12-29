package main

import (
	"fmt"
	"time"
)

func test5() {
	var t1 time.Time = time.Now()
	fmt.Printf("current time:%v\n", t1)
	fmt.Printf("%T\n", t1)
	fmt.Println(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second())

	// Unix time
	u_time := t1.Unix()
	ut1 := t1.UnixMilli()
	ut2 := t1.UnixMicro()
	ut3 := t1.UnixNano()
	fmt.Println(u_time, ut1, ut2, ut3)

	// 时间间隔
	//const (
	//	Nanosecond  Duration = 1
	//	Microsecond          = 1000 * Nanosecond
	//	Millisecond          = 1000 * Microsecond
	//	Second               = 1000 * Millisecond
	//	Minute               = 60 * Second
	//	Hour                 = 60 * Minute
	//)
	// time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。
	// time.Duration表示1纳秒，time.Second表示1秒。
	now := time.Now()
	t := now.Add(time.Hour)
	t2 := now.Add(time.Duration(1 * 60 * 60 * 1000 * 1000 * 1000))
	fmt.Println(t)
	fmt.Println(t2)
	time.Sleep(time.Second)
	fmt.Println(time.Now().Sub(now))
	fmt.Println(time.Now().Before(now))
	fmt.Println(time.Now().After(now))

	// 定时器? Tick
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))

	// 小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数
	fmt.Println(now.Format("2006??01??02 15:04:05.00000")) // 2022/02/27 00:10:42.960
	// 小数点后写9，会省略末尾可能出现的0
	fmt.Println(now.Format("2006/01/02 15:04:05.99999999999999")) // 2022/02/27 00:10:42.96

	// 只格式化时分秒部分
	fmt.Println(now.Format("15:04:05"))
	// 只格式化日期部分
	fmt.Println(now.Format("2006.01.02"))

	// 对于从文本的时间表示中解析出时间对象，time包中提供了time.Parse和time.ParseInLocation两个函数。
	// 其中time.Parse在解析时不需要额外指定时区信息。
	t3, err := time.Parse("2006?01?02 15:04:05.99999", "2023?05?05 1:02:03") // 必须有前置的0
	if err != nil {
		fmt.Printf("time parse failed, %s", err.Error())
		return
	}
	fmt.Println(t3)
}
func main() {
	test5()
}
