package sync

import (
	"fmt"
	"sync"
)

// 使用sync.WaitGroup结构体替代time.Sleep()实现多线程同步。

func WaitGroupExample() {
	var wait sync.WaitGroup
	var mainWait sync.WaitGroup

	// 循环创建协程，并使用WaitGroup实现同步
	mainWait.Add(10)
	fmt.Println("开始执行")
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func(i int) {
			fmt.Println("协程", i, "开始执行")
			//... 协程执行代码
			// 执行结束，调用Done
			wait.Done()
			mainWait.Done()
		}(i)
		// 主协程等待执行完毕
		wait.Wait()
	}
	// 确定所有协程都执行完毕
	mainWait.Wait()
	fmt.Println("所有协程执行完毕")
}
