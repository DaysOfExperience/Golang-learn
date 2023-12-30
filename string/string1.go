package main

import "fmt"

func test1() {
	s1 := "Hello Golang"
	// 字符串每个字符对应1个整数 (其它编码可能不止1个)
	for _, ch := range s1 {
		fmt.Printf("[%d] - %v\n", ch, ch)
	}
}

func test2() {
	s1 := "abcdef"
	var arr = []byte(s1)
	arr[0] = 65
	s1 = string(arr) // []byte转化回string
	fmt.Println(s1)
}

func test3() {
	s := "重估一切价值"
	arr := []byte(s)
	fmt.Println(len(arr)) // 几个字节
	fmt.Println(arr)
	arr[0] = 64
	fmt.Println(arr)
	s = string(arr)
	fmt.Println(s)
	// 18
	//[233 135 141 228 188 176 228 184 128 229 136 135 228 187 183 229 128 188]
	//[64 135 141 228 188 176 228 184 128 229 136 135 228 187 183 229 128 188]
	//@��估一切价值
}

func main() {
	//test1()
	//test2()
	test3()
}
