package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Goroutine 是 Go 程序中最基本的并发执行单元。
// Go运行时会智能地将 m个goroutine 合理地分配给n个操作系统线程，实现类似m:n的调度机制，不再需要Go开发者自行在代码层面维护一个线程池。
// 在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能——goroutine
// 当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个 goroutine 去执行这个函数就可以了，就是这么简单粗暴。

// Go语言中使用 goroutine 非常简单，只需要在函数或方法调用前加上go关键字就可以创建一个 goroutine ，从而让该函数或方法在新创建的 goroutine 中执行。
var wg sync.WaitGroup

func print() {
	fmt.Println("Just do it")
	wg.Done()
}
func test0() {
	wg.Add(1)
	go print()
	fmt.Println("main func printing")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) { // 这里不要实现闭包
			fmt.Println(num)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func func1() {
	for i := 0; i < 10000; i++ {
		fmt.Println("func1", i)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 10000; i++ {
		fmt.Println("func2", i)
	}
	wg.Done()
}
func test2() {
	runtime.GOMAXPROCS(1) // 即使是核心数为1, 也是并发, fun1 fun2也不是串型的
	wg.Add(2)
	go func1()
	go func2()
	wg.Wait()
}
func main() {
	test2()
}
