package main

import (
	"fmt"
	"strings"
)

// 函数 - 匿名函数与闭包

func func9() {
	// 1. 匿名函数保存在变量中
	fun := func(a int, b int) int {
		return a + b
	}
	ret := fun(1, 2)
	fmt.Println(ret)

	// 匿名函数直接调用
	func() {
		fmt.Println("重估一切价值")
	}()
}

// 使用场景: 匿名函数多用于实现闭包
// 闭包 = 函数 + 外层变量的引用

// 闭包的概念是指一个函数和与其相关的引用环境组合的一个整体（即闭合的）。闭包让你可以在一个函数内部访问到其外部作用域的变量。
// 在 Go 中，匿名函数是实现闭包的一种方式。闭包是一个函数，它“捕获”了在其外部作用域定义的一个或多个变量的引用。

// 闭包示例1:
// 此函数返回了一个闭包
// 闭包使得 adder 函数可以生成多个独立的累加器，每个累加器都有自己的私有状态。

// 疑惑: 为什么局部变量x在adder执行结束后, 在后续的闭包执行时, x可以一直维持这个累加状态?
// gpt4: 您的疑问非常合理，这涉及到闭包的一个核心特性：闭包能够捕获并持有其外部作用域中的变量，即使外部函数的执行已经完成。
// 在 Go 语言中，当一个函数（如 adder）返回一个闭包时，这个闭包会记住并持有它所捕获的变量（如 x）的引用。
// 持久性状态：即使 adder 函数执行完毕，局部变量 x 并不会像通常的局部变量那样消失。
// 相反，由于闭包仍在引用 x，因此 x 的生命周期会延续，直到闭包本身不再被使用。
// 封闭作用域：x 对于 adder 函数外的其他部分是不可见的。它只在闭包内部可见且可修改。
// 这提供了一种封装变量的方式，使得变量的作用域限定在闭包内。

// 每次调用 adder 函数时，它都会创建一个新的局部变量 x 和一个新的闭包。 也就是x与闭包是一对一的
// 这个闭包持有对其自己版本的 x 的引用。每次调用闭包时，都会修改它所持有的那个特定 x 的值。
// 这就是为什么即使 adder 函数的执行已经完成，闭包仍然能够访问并修改 x 的原因。
func adder() func(int) int {
	x := 0
	// 这个匿名函数是一个闭包，因为它访问了定义在 adder 函数外部的变量 x
	return func(num int) int {
		x += num // x: 外层变量的引用
		return x
	}
}
func test_adder() {
	fun := adder()
	fmt.Println(fun(10))
	fmt.Println(fun(20))
	fmt.Println(fun(20))
}

// 闭包示例2: 利用闭包实现了一个文件后缀检查器, 若后缀符合要求则没事, 不符合则加后缀
func makeSuffixFunc(suffix string) func(string) string {
	return func(s string) string {
		if strings.HasSuffix(s, suffix) == true {
			return s
		} else {
			return s + suffix
		}
	}
}
func test_suffix() {
	txtSuffix := makeSuffixFunc(".txt")
	name := "name"
	name = txtSuffix(name)
	fmt.Println(name)

	hello := txtSuffix("hello.txt")
	fmt.Println(hello)

	exeSuffix := makeSuffixFunc(".exe")
	file := exeSuffix("golang")
	fmt.Println(file)
}
func main() {
	// test_adder()
	// test_suffix()
}

// panic / recover
// 略了
