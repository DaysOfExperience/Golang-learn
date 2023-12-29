package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 一个例子: 主goroutine如何优雅地结束掉子goroutine
// 1. 用一个全局的exit bool变量
// 2. 用一个channel也可以
var wg sync.WaitGroup

// 一个context的使用示例
func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func test() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}

// context的Value方法如何使用
// 初步了解Value方法的使用: 其实就是把context看作一个map[interface{}]interface{}
func step1(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "key1", "value1")
	return child
}
func step2(ctx context.Context) context.Context {
	child := context.WithValue(ctx, 11, 11.11)
	return child
}
func step3(ctx context.Context) {
	// 把ctx的key value都打印出来
	fmt.Println("key1", ctx.Value("key1"))
	fmt.Println(11, ctx.Value(11))
}
func test1() {
	var ctx context.Context = context.TODO() // ctx就是一个最原始的context的实现
	ctx2 := step1(ctx)
	ctx3 := step2(ctx2)
	step3(ctx3)
}
func test2() {
	ctx, cancelFunc := context.WithTimeout(context.TODO(), time.Second) // 1s后自动超时
	defer cancelFunc()
	select {
	case <-ctx.Done(): // 返回一个只读的管道
		// 当超时后, deadline自动被调用, 此管道被关闭, 就可以读取了
		err := ctx.Err() // 此时Err返回的error会有关闭的原因
		fmt.Println(err)
	}
}

// 验证父子context不同的过期时间, 一定是以最先到期的那个时间为准的
func test3() {
	ctx1, cancelFunc := context.WithTimeout(context.TODO(), time.Second)
	defer cancelFunc()
	t1 := time.Now()
	ctx2, cancelFunc2 := context.WithTimeout(ctx1, time.Millisecond*500) // 500ms
	defer cancelFunc2()
	select {
	case <-ctx2.Done():
		fmt.Println(ctx2.Err(), time.Now().Sub(t1))
	}
}

// 之前是超时时间到了, 自动调用deadline关闭管道
// 而现在是主动调用cancel关闭管道
// cancel
func test4() {
	cxt, cancel := context.WithCancel(context.TODO())
	t1 := time.Now()
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()
	select {
	case <-cxt.Done():
		fmt.Println(time.Now().Sub(t1))
		fmt.Println(cxt.Err())
	}
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "exit")
			fmt.Println(name, ctx.Err())
			return
		default:
			fmt.Println(name, "watching...")
			time.Sleep(time.Second)
		}
	}
}

// xly
func test_cancel() {
	ctx, cancel := context.WithCancel(context.TODO())
	go watch(ctx, "goroutine1")
	go watch(ctx, "goroutine2")
	time.Sleep(time.Second * 6)
	fmt.Println("end watching")
	cancel()
	time.Sleep(time.Second)
}

//var wg sync.WaitGroup

func watch2(ctx context.Context, name string, time2 time.Time) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "exit", time.Now().Sub(time2))
			fmt.Println(name, ctx.Err())
			wg.Done()
			return
		default:
			fmt.Println(name, "watching...")
			time.Sleep(time.Second)
		}
	}
}
func test_deadline() {
	ctx1, cancel1 := context.WithDeadline(context.TODO(), time.Now().Add(time.Second*5))
	ctx2, cancel2 := context.WithDeadline(context.TODO(), time.Now().Add(time.Second*5))
	begin := time.Now()
	defer cancel1()
	wg.Add(2)
	go watch2(ctx1, "goroutine 1", begin)
	go watch2(ctx2, "goroutine 2", begin)
	time.Sleep(time.Second * 3)
	cancel2()
	wg.Wait()
}
func main() {
	//test()
	//test1()
	//test2()
	//test3()
	//test4()
	//test_cancel()
	test_deadline()
}
