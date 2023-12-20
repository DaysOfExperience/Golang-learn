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
	var m = map[string]int{"hhh": 1, "zzz": 2} // 引用类型也可以有指针
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
	// *int, 10
	// *float64, 3.14
	// *bool, true
	// *string, "111"
	// *map[string]int, map[string]int{"hhh":1, "zzz":2}
	// *[]int, []int{1, 2, 3}
}

// invalid  int: 值类型
func changeInt1(num int) {
	num = 0
}
func changeInt(num *int) {
	*num = 0
}

// valid resultful effective
// map: 引用类型
func changeMap(m map[int]int) {
	m[2] = 2
}
func changeMap2(m *map[int]int) {
	(*m)[3] = 3
	// m[3] = 3   // error
}

func ptr2() {
	// 值类型与引用类型
	num := 10
	m := map[int]int{1: 1}

	numPtr := &num
	mPtr := &m

	fmt.Println(num)
	changeInt1(num) // no
	fmt.Println(num)
	changeInt(numPtr) // yes
	fmt.Println(num)

	fmt.Println(m)
	changeMap(m) // yes
	fmt.Println(m)
	changeMap2(mPtr) // yes
	fmt.Println(m)
}

// new & make
// func new(Type) *Type
// 使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。
func test_new() {
	p1 := new(bool)        // false
	p2 := new(string)      // ""
	p3 := new([]int)       // nil
	p4 := new(map[int]int) // nil
	fmt.Println(*p1, *p2, *p3 == nil, *p4 == nil)
}

// make
// func make(t Type, size ...IntegerType) Type
// make也是用于内存分配的，区别于new，它只用于slice、map以及channel的内存创建，而且它返回的类型就是这三个类型本身，
// 而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
func test_make() {
	var m map[string]int // nil
	fmt.Println(m == nil)
	m = make(map[string]int, 10) // 分配空间

	var sli []string // nil
	fmt.Println(sli == nil)
	sli = make([]string, 0, 100) // 分配空间
}

// new与make的区别
// 1. 二者都是用来做内存分配的。
// 2. make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
// 3. 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
func main() {
	// ptr1()
	// ptr2()
	// test_new()
	test_make()
}
