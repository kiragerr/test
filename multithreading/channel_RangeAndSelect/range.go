package channel_test

import (
	"fmt"
	"time"
)

// 使用range遍历channel

func RangeChannel() {
	ch := make(chan int, 3)
	// 子协程写入
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("send", i)
		}
		// 确认写入完毕，关闭管道

		// 在写入端关闭管道，防止死锁。
		close(ch)
	}()

	// 主协程遍历
	for n := range ch {
		// for range 遍历管道，只有一个返回值。
		// 当管道关闭时，读取到空缓冲区不会阻塞，而是直接结束。

		fmt.Println("recieved", n)
		// 模拟数据处理:延迟 10ms
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Println("range channel done")
}
