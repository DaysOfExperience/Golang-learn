package main

import (
	"fmt"
)

func add_plus() (int, int) {
	return 1, 2
}

var Age = 11
var Salary = 12

// Age2 := 13     // error

func main() {
	// 变量的标准声明格式
	var name string
	var age int
	var job string
	var dict map[string]int
	fmt.Println(name+"11", age, job+"22", dict)
	// 批量声明
	var (
		a string = "hehe"
		b int    = 11
		c int64
		d float32
		e map[int]int
	)
	fmt.Println(a, b, c, d, e == nil)
	// 声明变量同时指定其初始值
	var name1 string = "hhh"
	var age1 int = 18

	// 省略类型声明, 让编译器根据变量初始值推导其类型
	var name2 = "zzz"
	var age2 = 19.01
	fmt.Println(name1, age1, name2, age2)

	// 短变量声明, 最常用的, 也就是上方这个的进一步简化
	// 函数外的每个语句都必须以关键字开始（var、const、func等） 因此:=不能使用在函数外。
	// 只能用于函数内
	name3 := "son of a bitch"
	age3 := 19
	fmt.Println(name3, age3)
	// 匿名变量, 比如函数返回值有两个, 不想接收某一个, 就用_匿名变量来接收
	// 匿名变量用一个下划线_表示
	// 匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。
	num1, _ := add_plus()
	fmt.Println(num1)

	fmt.Println("-----------------------------------------")
	const_test()
	fmt.Println("-----------------------------------------")
	iota_test()
}

func const_test() {
	const name string = "hhh"
	const age int = 18
	// const age2 := 3   //  error
	const (
		age2 int = 19
		age3     = 17
		age4     // 17
		age5     // 17
	)
	fmt.Println(name, age, age2, age3, age4, age5)
}

// iota是go语言的常量计数器，只能在常量的表达式中使用。
// iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解
// 为const语句块中的"行索引")。 使用iota能简化定义，在定义枚举时很有用。

// 也就是你在这个const这个批量声明里面是第几行
func iota_test() {
	const (
		a, b = iota, iota
		c    = iota
		d    = 1999
		e    = iota
	)
	fmt.Println(a, b, c, d, e)
}
