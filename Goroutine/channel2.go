package main

import (
	"fmt"
	"sync"
)

// 并发安全和锁

var (
	x   int64
	wg3 sync.WaitGroup
	m   sync.Mutex
)

// 不加保护的情况下, 两个goroutine并发访问共享资源, 会出现竞态问题
// 这两个 goroutine 在访问和修改全局的x变量时就会存在数据竞争，某个 goroutine 中对全局变量x的修改可能会
// 覆盖掉另一个 goroutine 中的操作，所以导致最后的结果与预期不符。

// 互斥锁是一种常用的控制共享资源访问的方法，它能够保证同一时间只有一个 goroutine 可以访问共享资源。Go 语言中使用sync包中提供的Mutex类型来实现互斥锁。
func fun() {
	for i := 0; i < 1000; i++ {
		// 使用互斥锁能够保证同一时间有且只有一个 goroutine 进入临界区，其他的 goroutine 则在等待锁；
		// 当互斥锁释放后，等待的 goroutine 才可以获取锁进入临界区，多个 goroutine 同时等待一个锁时，唤醒的策略是随机的。
		m.Lock()
		x = x + 1 // 临界区
		m.Unlock()
	}
	wg3.Done()
}

func test8() {
	wg3.Add(2)
	go fun()
	go fun()
	wg3.Wait()
	fmt.Println(x)
}
func main() {
	test8()
}
