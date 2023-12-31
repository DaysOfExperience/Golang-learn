package main

import "fmt"

// 其实切片就像是一个原生数组的指针引用一样, 且它可以引用数组的一部分区域
// 所以如果用一个数组构造一个切片, 对切片修改就是对数组的修改
// 又因为切片可以自动扩容, 所以如果切片扩容, 他就会指向新的区域(可能是新的地址, 可能不是), 而原来的数组大小和内容并不变
func main() {
	//Slice6()
	// Slice4()
	// Slice5()
	// Slice7()
	//Slice8()
	//Slice9()
	Slice0()
	//Slice1()
	//Slice2()
}

func Slice1() {
	// 切片的定义
	var sli1 []int               // 声明切片 nil
	var sli2 = []int{1, 2, 3, 4} // 声明并初始化切片 非nil
	var sli3 = []int{}           // 声明并初始化切片 非nil
	// 切片只能和nil比较, 因为是引用类型
	fmt.Println(sli1, sli2, sli3)
	fmt.Println(sli1 == nil, sli2 == nil, sli3 == nil)

	// 根据数组定义切片
	arr := [...]int{1, 2, 3, 4, 5}
	sli4 := arr[0:3]         // 根据数组定义切片
	fmt.Println(sli4)        // 1 2 3
	fmt.Printf("%T\n", sli4) // []int 类型

	// 切片再切片
	sli5 := sli4[:]          // sli4的全部
	fmt.Printf("%T\n", sli5) // []int 类型

	// make函数构造切片
	sli6 := make([]int, 3, 7)
	fmt.Println(sli6)      // 0 0 0
	fmt.Println(len(sli6)) // 3
	fmt.Println(cap(sli6)) // 7

	// 对切片再执行切片表达式时（切片再切片），high的上限边界是切片的容量cap(a)，而不是长度。
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]  2 3
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4]                                                      // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2)) // 5 1 1 !!!!

	arr5 := [5]int{1, 2, 3, 4, 5}
	slice := arr5[:2:3]                        // a[low : high : max]
	fmt.Println(slice, len(slice), cap(slice)) // [1 2] 2 3

	// 切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和nil比较。
}

func Slice2() {
	// 探究切片的底层原理
	// 如果是用数组来构造切片, 则切片的cap就是后面
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[:]
	s2 := arr[1:3]
	fmt.Printf("%d, %d | %d, %d\n", len(s1), cap(s1), len(s2), cap(s2))

	var a []int            // nil, 0
	var b = []int{}        // 非nil, 0
	var c = make([]int, 0) // 非nil, 0
	fmt.Println(cap(a), cap(b), cap(c))
}

// 切片是引用, 类似于指针!!!!
func Slice4() {
	a := make([]int, 3)
	b := a
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)
}

func Slice5() {
	// 切片的遍历
	s := []int{1, 2, 3, 4}
	for i := 0; i < len(s); i++ {
		println(s[i])
	}
	fmt.Println()
	for _, num := range s {
		println(num)
	}
}

// 切片新增元素 append
func Slice6() {
	arr := [3]int{1, 2, 3}
	var sli = arr[1:2]
	fmt.Println(len(sli)) // 1
	fmt.Println(cap(sli)) // 2
	sli = append(sli, 4)
	fmt.Println(sli) // 24   也就是这个切片的修改直接改变了这个数组
	fmt.Println(arr) // 124
	sli = append(sli, 6, 7, 8, 9)
	fmt.Println(sli)      // 246789
	fmt.Println(len(sli)) // 6
	fmt.Println(cap(sli)) // 6
	fmt.Println(arr)      // 124!!!!!!
	sli[0] = 0
	fmt.Println(sli) // 046789
	fmt.Println(arr) // 124!!!!!!
}

func Slice7() {
	var sli []int // nil
	for i := 0; i <= 10; i++ {
		sli = append(sli, i)
		fmt.Printf("%v %d %d %p\n", sli, len(sli), cap(sli), sli)
	}
	arr := []int{1, 2, 3}
	sli = append(sli, arr...) // 切片追加到另一个切片后方
}

func Slice8() {
	// copy函数复制切片
	a := []int{1, 2, 3, 4}
	var b = make([]int, 4)
	copy(b, a) // copy执行之前dst切片必须有足够的空间容量容纳a的内容
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)

	var c = make([]int, 0)
	fmt.Println(c == nil)
	//[1 2 3 4]
	//[100 2 3 4]
	//false
}

// 从切片中删除元素
func Slice9() {
	// 利用copy
	a := []int{1, 2, 3, 4, 5}
	a = append(a[:1], a[2:]...)
	fmt.Println(a) // 1 3 4 5
	fmt.Println(len(a), cap(a))
}

// Slice0 说明, 一个切片引用一个数组时, 扩容之前, 也就是数组容量足够时, 对切片的修改就是对数组的修改, 因为他就是原生数组直接切了一部分
// 而扩容之后, 他就会在新的空间, 对切片的修改不会影响原生数组
func Slice0() {
	arr := [3]int{1, 2, 3}
	sli := arr[1:2] // [2]  len:1 cap:2
	fmt.Printf("%p\n", sli)
	sli = append(sli, 4)
	fmt.Println(arr)        // 124
	fmt.Println(sli)        // 24
	fmt.Printf("%p\n", sli) // 没变, 因为没扩容
	sli[1] = 10
	fmt.Println(arr) // 12 10
	fmt.Println(sli) // 2 10
	sli = append(sli, 6)
	fmt.Printf("%p\n", sli) // 变了, 因为扩容了, 指向新的空间
	sli[1] = 8
	fmt.Println(arr) // 12 10  // 没变, 因为修改前, 切片已经指向新的地址了
	fmt.Println(sli) // 2 8 6

	//0xc000012140
	//[1 2 4]
	//[2 4]
	//0xc000012140
	//[1 2 10]
	//[2 10]
	//0xc000014180
	//[1 2 10]
	//[2 8 6]
}
