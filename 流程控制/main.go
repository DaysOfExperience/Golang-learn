package main

import "fmt"

func getNum() int {
	return 80
}

func main() {
	// for_test()
	// for_range()
	switch_test()
}

func if_test() {
	if false {
		fmt.Println("fuck you")
	} else if true {
		fmt.Println("hehe")
	}

	if num := getNum(); num >= 0 {
		fmt.Println("num >= 0")
	} else {
		fmt.Println("num < 0")
	}
}

func for_test() {
	// 标准的for
	numArray := [7]int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(numArray); i++ {
		fmt.Printf("%d-", numArray[i])
	}
	fmt.Println()
	// 省略第一个表达式
	i := 0
	for ; i < len(numArray); i++ {
		fmt.Printf("%d ", numArray[i])
	}
	// 省略第一个和第三个, 就是一个while
	fmt.Println()
	num := 1
	for num <= 10 {
		fmt.Printf("%d", num)
		num++
	}
	fmt.Println()
	// while(true)
	for {
		fmt.Println(num)
		break
	}
}

// Go语言中可以使用for range遍历数组、切片、字符串、map 及通道（channel）。 通过for range遍历的返回值有以下规律：
// 数组、切片、字符串返回索引和值。
// map返回键和值。
// 通道（channel）只返回通道内的值。
func for_range() {

}

func switch_test() {
	num := 3
	switch num {
	case 2:
		fmt.Println("num == 2")
	case 3:
		fmt.Println("num == 3")
	}
	switch n := 0; n {
	case 0:
		fmt.Println("0")
	case 1, 2, 3, 4, 5, 6:
		fmt.Println("1, 2, 3, 4, 5, 6")
	}
	switch {
	case num >= 0:
		fmt.Println("num >= 0")
		fallthrough // 不合适
	case num < 0:
		fmt.Println("num < 0")
	}
}
