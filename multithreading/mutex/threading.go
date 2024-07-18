package mutex

import (
	"fmt"
	"time"
)

// 理解管道的原理以及使用方法

// p1: 无缓冲管道的特点
// 无缓冲管道无缓冲区，无法储存数据，在同步过程使用会导致死锁

func NonBufferedChannelExample() {
	// 创建无缓冲区管道:
	ch := make(chan int)
	defer close(ch)

	// 创建子协程传入数据(防止死锁)

	go func() {
		ch <- 1
	}() // 立即执行

	// 主协程等待通道数据
	n := <-ch

	fmt.Println("Received data:", n)
}

// p2: 有缓冲管道:
// 当缓冲区有空余位置时，允许直接读取和写入
// 当缓冲区已满，若无协程进行读取，则写入阻塞,直到读出数据，才允许写入
// 当缓冲区为空，若无协程进行写入，则读取阻塞..
func BufferedChannelExample() {
	// 创建3个协程，分别用于储存数据，输出完成信号
	chForData := make(chan int, 3) // 缓冲区大小为2(单位为数据个数)
	// 用于子协程传出信号
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	// 关闭通道
	defer close(chForData)
	defer close(ch1)
	defer close(ch2)

	// 子协程一:写入到管道
	go func() {
		for i := 0; i < 10; i++ {
			chForData <- i // 写入数据
			fmt.Println("Data sent:", i)
		}
		// 当写入完毕，传出信号给主协程
		ch1 <- true
	}()

	// 子协程二:读取管道数据
	go func() {
		for i := 0; i < 10; i++ {
			// 读取延迟1毫秒
			time.Sleep(time.Millisecond)

			n := <-chForData // 读取数据
			fmt.Println("Data received:", n)
		}
		// 当读取完毕，传出信号给主协程
		ch2 <- true
	}()

	// 主协程等待子协程完毕:同步进程等待无缓冲管道数据流动

	fmt.Println("Data sent finished:", <-ch1)

	fmt.Println("Data received finished:", <-ch2)

	fmt.Println("All data sent and received")
}

// 管道注意点：
// 当管道值为nil时，无论如何读写都会导致阻塞(等待cpu允许执行)
// 管道值为nil的情况: 只声明管道 var chName chan T,管道值为nil(未分配内存)

// 关闭一个nil管道会导致panic

// 向关闭的管道写入或者读取也会导致panic

// 关闭已经关闭的管道也会导致panic
