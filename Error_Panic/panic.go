package main

import "fmt"

func test4() {
	fmt.Println("test4 before")
	panic("test panic")
	fmt.Println("Unreachable code")
}
func test5() {
	defer func() {
		if err := recover(); err != nil {
			// 发生了panic, 且recover生效了
			fmt.Println(err)
			fmt.Println("This is valid")
		}
	}()
	fmt.Println("test5 before")
	test4()
	fmt.Println("unreachable")
}
func test6() {
	fmt.Println("test6 before")
	test5()
	fmt.Println("test6 after")
}

// 函数发生了panic之后会一直向上传递, 如果直至main函数都没有recover()，程序将终止，如果是碰见了recover()，将被recover捕获。
// recover()只能恢复当前函数级或以当前函数为首的调用链中的函数中的panic(), 恢复后调用当前函数结束(test5), 但是调用此函数的函数继续执行(test6)
func main() {
	test6()
}
