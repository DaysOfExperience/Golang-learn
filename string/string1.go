package main

import "fmt"

func test1() {
	s1 := "Hello Golang"
	// 字符串每个字符对应1个整数 (其它编码可能不止1个)
	for _, ch := range s1 {
		fmt.Printf("[%d] - %v\n", ch, ch)
	}
	//var s2
}

func main() {
	test1()
}
