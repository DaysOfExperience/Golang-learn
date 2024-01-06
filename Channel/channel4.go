package main

import (
	"fmt"
	"time"
)

func fun1() {
	ch := make(chan int, 5)
	a := 1
	ch <- a
	fmt.Println("1 - ", a) // a1
	a = 2

	go func(c chan int) {
		num := <-c
		fmt.Println("2 - ", num) // a2
	}(ch)
	time.Sleep(time.Second)
	fmt.Println("3 - ", a) // a3

	//1 -  1
	//2 -  1
	//3 -  2
	// 说明是元素的拷贝
}

type people struct {
	name string
}

func fun2() {
	ch := make(chan *people, 5)
	p := &people{
		"aa",
	}
	fmt.Println("first - ", p.name) // aa

	ch <- p
	p.name = "bb"

	go func(c chan *people) {
		ptr := <-c
		fmt.Println("second - ", ptr.name) // bb
	}(ch)

	time.Sleep(time.Second)
	fmt.Println("third - ", p.name) // bb

	//first -  aa
	//second -  bb
	//third -  bb
}

func fun3() {
	ch := make(chan *people, 5)
	p := people{
		"aa",
	}
	ptr := &p
	ch <- ptr // 一个指针值拷贝进去
	fmt.Println("first - ", ptr)
	ptr = &people{"bb"} // 无效修改

	go func(c chan *people) {
		pp := <-c
		fmt.Println("second - ", pp)
	}(ch)

	time.Sleep(time.Second)
	fmt.Println("third - ", ptr) // 变化
}

// 4个goroutine, 编号1234, 每秒有一个goroutine打印出自己的编号, 1号打印1...按照12341234...打印出来
func test() {
	sli := make([]chan int, 4) // channel切片
	for i, _ := range sli {
		sli[i] = make(chan int)
		// ch是一个channel
		go func(i int) { // 闭包
			for {
				num := <-sli[i]
				fmt.Println(num)
				num++
				if num == 5 {
					num = 1
				}
				time.Sleep(time.Second)
				sli[(i+1)%4] <- num
			}
		}(i)
	}
	sli[0] <- 1
	select {}
}

func main() {
	//fun1()
	//fun2()
	//fun3()
	test()
}
