package channel_test

import (
	"fmt"
	"time"
)

// 使用select语句监听多通道。
// 使用Loop: for-select 循环监听通道
// 使用time.Afte()实现超时

func SelectExample() {
	// 创建数据通道，传递信息
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	ch3 := make(chan int, 2)
	defer func() {
		close(ch1)
		close(ch2)
		close(ch3)
	}()

	// 监听情况通信:
	l := make(chan int)

	// 子协程监听通道
	go func() {
		// Loop: for-seleclt 循环监听

		// 不使用default语句，使用timeAfter超时机制退出
	Loop:
		for {
			select {
			case n, ok := <-ch1:
				fmt.Println("ch1:", n, ok)
			case n, ok := <-ch2:
				fmt.Println("ch2:", n, ok)
			case n, ok := <-ch3:
				fmt.Println("ch3:", n, ok)
			// 实现超时
			// 当没有数据传入超过两秒，退出循环
			case <-time.After(time.Second * 2):
				fmt.Println("timeout")
				break Loop
			}
		}

		// 循环结束时，通知主协程
		l <- 1
	}()

	// 主协程传入数据
	go sent(ch1, "ch1")
	go sent(ch2, "ch2")
	go sent(ch3, "ch3")

	<-l // 等待子协程结束

	fmt.Println("select example end")
}

// 发送数据(单向管道)
func sent(ch chan<- int, chName string) {
	for i := 0; i < 4; i++ {
		ch <- i
		fmt.Println(chName, "data sent:", i)
	}
}
