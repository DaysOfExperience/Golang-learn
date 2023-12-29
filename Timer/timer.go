package main

import (
	"fmt"
	"sync"
	"time"
)

// 超时后, C channel会有数据写入
//type Timer struct {
//	C <-chan Time
//	r runtimeTimer
//}

// Timer: 可以定时, 可以暂停, 可以重新计时, 可以定时若干时长执行某个方法, 可以配合select使用
func test1() {
	// 创建一个Timer
	t1 := time.NewTimer(5 * time.Second) // *Timer
	begin := time.Now()
	time.Sleep(2 * time.Second)
	b1 := t1.Stop()
	fmt.Println(b1) // 未超时, Stop所以true
	t1.Reset(time.Second)
	<-t1.C
	fmt.Println(time.Now().Sub(begin))
	t1.Reset(time.Second * 2)
	<-t1.C
	fmt.Println(time.Now().Sub(begin))

	fmt.Println("--------")
	t2 := time.NewTimer(3 * time.Second)
	begin2 := time.Now()
	b2 := t2.Reset(time.Second) // 没有过期/暂停的timer也可以Reset
	fmt.Println(b2)
	<-t2.C
	fmt.Println(time.Now().Sub(begin2))
}
func test2() {
	//AfterFunc: 定时, 若干时间之后就执行某个函数
	fun := func() {
		fmt.Println("Hello")
	}
	var t *time.Timer = time.AfterFunc(time.Second*2, fun) // *Timer
	time.Sleep(time.Second * 3)
	b := t.Stop()
	fmt.Println(b)
}
func test3() {
	// After
	// 主要和select一起用, 比如一个case是从After返回的channel读取数据, 也就是若干时间的定时
	// 其它case可以是一些业务逻辑, 若若干时间依旧没有业务发生, 则After的channel将超时
	// Timer几乎是包含了time.After的作用
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "test"
	}()
	select {
	case s := <-ch:
		fmt.Println("get a value from ch: %v", s)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout...")
	}
}

// Ticker
func test4() {
	var wg sync.WaitGroup
	ticker := time.NewTicker(time.Second)
	ch := make(chan interface{})
	wg.Add(1)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println("tick once...")
			case <-ch:
				fmt.Println("exit...")
				wg.Done()
				return
			}
		}
	}()
	time.Sleep(3 * time.Second)
	ch <- 1
	close(ch)
	wg.Wait()
}
func main() {
	//test1()
	//test2()
	//test3()
	test4()
}
