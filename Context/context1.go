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

func test1() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}

func main() {
	test1()
}
