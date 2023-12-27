package main

import "fmt"

// 数组长度不可以变~ 类似于C++的原生数组
func main() {
	arr4()
}

func arr1() {
	var arr2 [4]int            // 定义一个数组, 自动初始化为默认值
	var arr1 = [3]int{1, 2, 3} // 定义并初始化全部元素
	var arr3 [3]int
	arr3 = arr1
	// arr2 = arr1     // arror
	fmt.Println(arr2)
	fmt.Println(arr3)

	var a2 [2]string                        // 两个空串, 定义, 初始化为默认值
	var a1 = [4]string{"111", "222", "333"} // 定义并初始化部分元素
	fmt.Println(len(a1))
	fmt.Println(a2)
	fmt.Println(a1)

	var a3 = [...]int{1, 2, 3, 4} // 自动推导数组元素个数
	fmt.Println(a3)
	fmt.Printf("%T\n", a3)

	var a4 = [...]int{0: 1, 5: 5} // 指定下标的元素, 其余未指定元素为默认值
	fmt.Println(a4)
}

// len: 数组: 数组元素个数 string: 字节数 切片: 元素个数
func arr2() {
	var arr = [5]int{0: 1, 4: 5}
	fmt.Println(len(arr))
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
	for index, num := range arr {
		fmt.Printf("%d:%d ", index, num)
	}
}

func arr3() {
	// 多维数组
	var arr = [3][4]int{
		{1, 2, 3, 4},
		{2, 3},
	}
	fmt.Println(arr)
	for _, a := range arr {
		for i, num := range a {
			fmt.Printf("%d:%d ", i, num)
		}
		fmt.Println()
	}

	// 多维数组只有第一层可以使用...来让编译器推导数组长度。
	var arr2 = [...][2]string{
		{"11", "22"},
		{"hh"},
	}
	for i := 0; i < len(arr2); i++ {
		for _, s := range arr2[i] {
			if s == "" {
				fmt.Printf("---\t")
			} else {
				fmt.Printf("%s\t", s)
			}
		}
		fmt.Println()
	}
}

// 参数是副本, 值类型
func test(a [3]int) {
	a[0] = 100
}
func arr4() {
	// 数组是值类型
	// 赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	// 数组的传参/赋值, 都是拷贝构造或拷贝赋值
	arr := [3]int{1, 2, 3}
	test(arr)
	fmt.Println(arr) // 123

	i := 10
	var ptr *int = &i
	var pptr **int = &ptr
	fmt.Println(ptr, pptr) // 两个指针
	var p *[3]int = &arr   // 数组指针
	arr2 := *p             // 数组, 拷贝构造
	fmt.Println(arr2)      // 123
	arr2[1] = 10
	fmt.Println(arr)  // 123
	fmt.Println(arr2) // 1103
	// 数组支持 “=="、”!=" 操作符，因为内存总是被初始化过的。
	fmt.Println(arr == arr2) // false
	arr2[1] = 2
	fmt.Println(arr == arr2) // true
}
