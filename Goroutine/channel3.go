package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁是完全互斥的，但是实际上有很多场景是 读多写少 的，当我们并发的去读取一个资源而不涉及资源修改的时候
// 是没有必要加互斥锁的，这种场景下使用读写锁是更好的一种选择。读写锁在 Go 语言中使用sync包中的RWMutex类型。

// 当一个 goroutine 获取读锁之后，其他的 goroutine 如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
// 当一个 goroutine 获取写锁之后，其他的 goroutine 无论是获取读锁还是写锁都会等待。

var (
	xx     int64
	wg4    sync.WaitGroup
	Lock   sync.Mutex
	RWLock sync.RWMutex
)

func read() {
	//Lock.Lock()
	RWLock.RLock()
	time.Sleep(time.Millisecond)
	//Lock.Unlock()
	RWLock.RUnlock()
	wg4.Done()
}

func write() {
	//Lock.Lock()
	RWLock.Lock()
	time.Sleep(10 * time.Millisecond)
	xx = xx + 1
	//Lock.Unlock()
	RWLock.Unlock()
	wg4.Done()
}

func test9() {
	start := time.Now()

	for i := 0; i < 1000; i++ {
		wg4.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg4.Add(1)
		go write()
	}
	wg4.Wait()
	fmt.Println(time.Now().Sub(start))
}

// sync.Once 确保某个函数在并发场景下只被安全地执行一次 例如某个配置文件的加载或者某大型map的加载

// map在并发场景下不安全
// fatal error: concurrent map writes错误。我们不能在多个 goroutine 中并发对内置的 map 进行读写操作，否则会存在数据竞争问题。
//var mm = make(map[int]int)

// Go语言的sync包中提供了一个开箱即用的并发安全版 map——sync.Map。开箱即用表示其不用像内置的 map 一样使用 make 函数初始化就能直接使用。
// var sync_m = sync.Map{}
var sync_m sync.Map

//	func get3(n int) (ret int) {
//		ret = mm[n]
//		return
//	}
//
//	func set3(k int, v int) {
//		mm[k] = v
//	}
func test10() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		// 每个并发函数都是, set一下, get一下
		// 并且它们的key是不同的
		go func(num int) {
			//set3(num, num+100)
			sync_m.Store(num, num+100)
			//fmt.Printf("set %d, get %d\n", num, get3(num))
			value, _ := sync_m.Load(num)
			fmt.Printf("set %v, get %v\n", num, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func main() {
	//test9()
	test10()
}
