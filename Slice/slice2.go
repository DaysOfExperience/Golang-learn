package main

import "fmt"

func test1() {
	s1 := make([]int, 0, 4)
	s1 = append(s1, 1)
	s2 := append(s1, 2)
	s3 := append(s1, 3)
	fmt.Printf("%v, address:%p, %p\n", s1, &s1, &s1[0])
	fmt.Printf("%v, address:%p, %p\n", s2, &s2, &s2[0])
	fmt.Printf("%v, address:%p, %p\n", s3, &s3, &s3[0])
	// s3的append对s2之前的append产生了覆盖, 底层都是一个数组
	//[1], address:0xc000008048, 0xc000014140
	//[1 3], address:0xc000008060, 0xc000014140
	//[1 3], address:0xc000008078, 0xc000014140
	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	//[100]
	//[100 3]
	//[100 3]

	//num := s1[1]
	//fmt.Println(num)
	// panic: runtime error: index out of range [1] with length 1
	s4 := append(s1, 2, 3, 4, 5) // 超过了容量4
	// 此时, s4因为append之后的数据完全超过了原始数组空间的容量
	// 所以在新的地方扩容
	fmt.Printf("%v, address:%p, %p\n", s1, &s1, &s1[0])
	fmt.Printf("%v, address:%p, %p\n", s2, &s2, &s2[0])
	fmt.Printf("%v, address:%p, %p\n", s3, &s3, &s3[0])
	fmt.Printf("%v, address:%p, %p\n", s4, &s4, &s4[0])
	//[100], address:0xc000094030, 0xc000092040
	//[100 3], address:0xc000094048, 0xc000092040
	//[100 3], address:0xc000094060, 0xc000092040
	//[100 2 3 4 5], address:0xc000094108, 0xc0000ac080
	fmt.Println(len(s1), cap(s1)) // 1 4
	fmt.Println(len(s2), cap(s2)) // 2 4
	fmt.Println(len(s3), cap(s3)) // 2 4
	fmt.Println(len(s4), cap(s4)) // 5 8  (可能是2倍扩容
}
func test2() {
	s1 := make([]int, 4, 4)
	s2 := append(s1, 5)
	fmt.Printf("%v, address:%p, %p\n", s1, &s1, &s1[0])
	fmt.Printf("%v, address:%p, %p\n", s2, &s2, &s2[0])
	//[0 0 0 0], address:0xc000008048, 0xc000014140
	//[0 0 0 0 5], address:0xc000008060, 0xc000010240
}

func main() {
	test1()
	//test2()
}
