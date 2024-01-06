package main

import (
	"fmt"
	"sync"
	"time"
)

func test1() {
	var ch1 chan int         // channel是引用类型, 零值:nil
	fmt.Println(ch1 == nil)  // true
	ch1 = make(chan int, 10) // 最多有10个, 也就是生产消费模型中的交易场所最多存10个元素
	ch1 <- 10                // 往channel中发送值
	num := <-ch1             // 从channel中接收值
	fmt.Println(num)
	close(ch1) // 关闭channel
	// 一个通道值是可以被垃圾回收掉的。通道通常由发送方执行关闭操作，并且只有在接收方明确等待通道关闭的信号时才需要执行关闭操作。
	// 它和关闭文件不一样，通常在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。
}

// 对一个关闭的通道再发送值就会导致 panic。
// 对一个关闭的通道进行接收会一直获取值直到通道为空。
// 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
// 关闭一个已经关闭的通道会导致 panic。

func get(ch chan string, group *sync.WaitGroup) {
	s := <-ch
	fmt.Println(s)
	group.Done()
}
func put(ch chan string, group *sync.WaitGroup) {
	ch <- "Just Do It"
	close(ch)
	group.Done()
}

// 使用无缓冲通道进行通信将导致发送和接收的 goroutine 同步化。因此，无缓冲通道也被称为同步通道。
func test3() {
	ch := make(chan string) //无缓冲通道, 必须有一个接收者才能发送成功, 否则阻塞. 必须有一个发送者才能接收成功, 否则阻塞
	wg := sync.WaitGroup{}
	wg.Add(2)
	go get(ch, &wg)
	go put(ch, &wg)
	wg.Wait()
}

// 只要通道的容量大于零，那么该通道就属于有缓冲的通道，通道的容量表示通道中最大能存放的元素数量。
// 当通道内已有元素数达到最大容量后，再向通道执行发送操作就会阻塞，除非有从通道执行接收操作。
// 我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。
var wg2 sync.WaitGroup

func put2(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg2.Done()
}
func get2(ch chan int) {
	time.Sleep(time.Second)
	// 此时chan已关闭, 获取值方法1: 双返回值来获取
	//for {
	//	v, ok := <-ch
	//	if ok {
	//		fmt.Printf("%v ", v)
	//	} else {
	//		// 没有元素了
	//		fmt.Println("chan is empty and closed", v)
	//		break
	//	}
	//}

	// 2. for range获取
	// 通常我们会选择使用for range循环从通道中接收值，当通道被关闭后，会在通道内的所有值被接收完毕后会自动退出循环。
	for v := range ch {
		fmt.Printf("%v ", v)
	}
	fmt.Println("chan is empty and closed")
	wg2.Done()
	// 注意: 目前Go语言中并没有提供一个不对通道进行读取操作就能判断通道是否被关闭的方法。
}
func test4() {
	// 有缓冲通道chan, 其实就是一个有若干存储空间的生产消费模型中的交易场所
	ch := make(chan int, 20)
	wg2.Add(2)
	go put2(ch)
	go get2(ch)
	wg2.Wait()
}

// 单向通道 其实很简单, 就类似于C++中的const成员函数一样
// 只是在传递chan类型变量时, 限制这个channel在该函数中只能发/收, 更安全一点

// <- chan int // 只接收通道，只能接收不能发送
// chan <- int // 只发送通道，只能发送不能接收
// 对于只接收的通道进行close是不允许的, 因为默认通道的关闭操作应该由发送方来完成。
//var wg3 sync.WaitGroup   // 没必要

func producer() <-chan int {
	ch := make(chan int, 2)
	//wg3.Add(1)
	// 闭包
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()
	//wg3.Done()
	return ch
}
func consumer(ch <-chan int) {
	// ch是一个只能接收, 不能发送的channel
	for v := range ch {
		fmt.Printf("%v ", v)
	}
	fmt.Println("channel is closed")
}
func test5() {
	ch := producer()
	consumer(ch)
	//wg3.Wait()
}

// worker pool: goroutine池
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker#%d gets job#%d\n", id, job)
		results <- job * 2
		time.Sleep(time.Millisecond * 500)
	}
}
func test6() {
	jobs := make(chan int, 100)
	results := make(chan int, 200)

	for j := 0; j < 5; j++ {
		// 开启5个goroutine
		go worker(j, jobs, results)
	}
	for i := 0; i < 100; i++ {
		// 发送100个任务到任务池中
		jobs <- i
	}
	close(jobs) // 发送完毕close channel
	// 输出这些任务的执行结果
	for i := 0; i < 100; i++ {
		ret := <-results // 取出结果
		fmt.Printf("%v\n", ret)
	}
	//close(results) // 不能for range 因为生产者那边没有进行close
}

// select

// 其实就是把一个channel当作一个套接字, 和网络的IO多路复用非常像!!!!
// 可处理一个或多个 channel 的发送/接收操作。
// 如果多个 case 同时满足，select 会随机选择一个执行。
// 对于没有 case 的 select 会一直阻塞，可用于阻塞 main 函数，防止退出。
// 每个 case 分支会对应一个通道的通信（接收或发送）过程。
func test7() {
	//ch := make(chan int, 1)
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		select {
		case num := <-ch:
			fmt.Printf("get num:%d\n", num)
		case ch <- i:
		}
	}
}
func main() {
	test6()
}
