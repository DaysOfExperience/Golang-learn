package main

import "fmt"

// Go语言中的指针不能进行偏移和运算，是安全指针。
// Go语言中的指针, 3个概念：指针地址、指针类型和指针取值。
// Go语言中的指针不能进行偏移和运算，因此Go语言中的指针操作非常简单，我们只需要记住两个符号：&（取地址）和*（根据地址取值）。
func ptr1() {
	var i int = 10
	var f float64 = 3.14
	var b bool = true
	var s string = "111"
	var m = map[string]int{"hhh": 1, "zzz": 2}
	var sli = []int{1, 2, 3}

	pi := &i
	pf := &f
	pb := &b
	ps := &s
	pm := &m
	psli := &sli

	fmt.Printf("%T, %#v\n", pi, *pi)
	fmt.Printf("%T, %#v\n", pf, *pf)
	fmt.Printf("%T, %#v\n", pb, *pb)
	fmt.Printf("%T, %#v\n", ps, *ps)
	fmt.Printf("%T, %#v\n", pm, *pm)
	fmt.Printf("%T, %#v\n", psli, *psli)
}
func main() {
	ptr1()
}
