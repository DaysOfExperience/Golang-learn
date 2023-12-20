package main

import "fmt"

// 两个命名的返回值可以直接使用, 不需要再定义, return时也可以省略返回的变量
func func1(a, b int, name ...string) (add, sub int) {
	fmt.Printf("type of ...string : %T\n", name)
	// 可变参数实际上就是一个切片类型
	for _, s := range name {
		fmt.Printf("%v ", s)
	}
	fmt.Println()
	add = a + b
	sub = a - b
	return
}

func func2(a int, b int, name ...string) (int, int) {
	fmt.Printf("type of ...string : %T\n", name)
	// 可变参数实际上就是一个切片类型
	for _, s := range name {
		fmt.Printf("%v ", s)
	}
	fmt.Println()
	add := a + b
	sub := a - b
	return add, sub
}
func func3() []int {
	return nil // nil可以作为切片的返回值
}
func func_end() {
	name1 := "hhhh"
	name2 := "zzzz"
	ret1, ret2 := func1(1, 2, name1, name2)
	fmt.Println(ret1, ret2)
	ret1, ret2 = func2(1, 2, name1, name2)
	fmt.Println(ret1, ret2)
}

// 全局变量, 局部变量, 局部和全局同名优先访问局部
// for语句/函数内的代码块内定义的变量只在代码块内有效
// 这些和C++是一样的

// ====================================================================

// 函数可以作为变量, 参数, 返回值
func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func div(a, b int) int {
	return a / b
}

func mod(a, b int) int {
	return a % b
}

func calculate(a, b int, cal func(int, int) int) int {
	// cal就是一个函数类型的参数
	return cal(a, b)
}

func func4() {
	a := 1
	b := 2
	fmt.Println(calculate(a, b, add))
	fmt.Println(calculate(a, b, sub))
	fmt.Println(calculate(a, b, div))
	fmt.Println(calculate(a, b, mod))

	fun1 := calculate
	fmt.Println(fun1(a, b, add))

	fun2 := test_return("+")
	fmt.Println(fun2(a, b))
}

// 和切片作为函数返回值一样, 函数作为函数返回值时也可以返回nil
func test_return(s string) func(int, int) int {
	if s == "+" {
		return add
	} else if s == "-" {
		return sub
	} else if s == "/" {
		return div
	} else {
		return mod
	}
}

// 使用type关键字定义一个函数类型
func func5() {
	type func_ func(int, int) int
	var function func_
	function = add
	fmt.Println(function(1, 2))
}

// defer
// Go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行
// 也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。

// 应用场景: 由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。
// defer时机:...

// （提示：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值, 注册之后参数就确定了!）
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func func6() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y)) // 这里: A 1 2 3
	x = 10
	defer calc("BB", x, calc("B", x, y)) // 这里: B 10 2 12
	y = 20
	// 返回之前: BB 10 12 22
	// AA 1 3 4
}

// A 1 2 3
// B 10 2 12
// BB 10 12 22
// AA 1 3 4

// 匿名函数: 和普通函数的定义的唯一区别就是 没有函数名
func fun3() func(int) { // 返回值: 一个函数, int参数且无返回值
	// 下面就是一个匿名函数
	// 匿名函数可以直接定义+执行, 也可以赋值给一个变量, 也就是函数变量
	fun := func() {
		fmt.Println("hhh")
	}
	fun() // 1. 把匿名函数赋值给变量, 然后执行
	func() {
		fmt.Println("uhhh")
	}() // 2. 匿名函数直接执行

	// 匿名函数也可以直接作为返回值返回
	return func(x int) {
		fmt.Println(x)
	}
}
func fun33() {
	f := fun3()
	f(2)

	type func_ = func(int, int) int
	var f_ func_ // nil
	fmt.Println(f_ == nil)
}
