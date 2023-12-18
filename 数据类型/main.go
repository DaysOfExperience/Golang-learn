package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	func1()

	fmt.Println("************************************")
	str()

	fmt.Println("************************************")
	character()
}
func func1() {
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxFloat64)
	b := true
	fmt.Println(b && false)
	s1 := "hhhhh\n"
	num1, _ := fmt.Printf(s1)
	fmt.Print(num1)
	s2 := "人生苦短go"
	fmt.Println(len(s2)) // 字节数
}
func str() {
	fmt.Println("hh" + "zz") // 字符串拼接
	// 字符串分割
	s2 := "ren sheng ku duan let's go"
	s3 := strings.Split(s2, " ")
	fmt.Println(s3)        // []string
	fmt.Printf("%T\n", s3) // 分割之后的类型是什么
	fmt.Println(strings.Contains(s2, "ku"))
	fmt.Println(strings.HasPrefix(s2, "ren"))
	fmt.Println(strings.HasSuffix(s2, "go"))
	fmt.Println(strings.HasSuffix(s2, "fuck"))
	fmt.Printf("\"ku\"在s2的下标为:%d\n", strings.Index(s2, "ku"))

	s4 := []string{"111", "222", "333"} // 切片
	fmt.Println(strings.Join(s4, "-"))
}

func character() {
	// byte uint8的别名  代表一个ASCII码字符。
	// rune int32的别名  代表一个 UTF-8字符。
	var c1 byte = '1'
	var c2 rune = '和'
	fmt.Printf("%T - %T\n", c1, c2)
	s := "人生苦短fuckyou" // 一个中文不是一个字节, 所以按照这样的是不行的
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()
	for i, c := range s {
		fmt.Printf("%c - %d\n", c, i)
	}
}
