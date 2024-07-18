package mutex

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 学习同步进程下用于控制协程间通信的三种方法:

// 1 互斥锁:

// 2 读写锁: RWMutex
func RWMutexTest() {
	var rw sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(12)
	// 互斥量:
	data := 0

	// 创建多个协程(读多写少)

	// 读协程:
	go func() {
		for i := 0; i < 7; i++ {
			go func() {
				// 随机读取时间
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))

				// 上读锁:
				rw.RLock()
				fmt.Println("取得读锁")

				// 模拟读取耗时
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100))) // 随机耗时

				fmt.Println("读取数据:", data)

				// 释放读锁:
				fmt.Println("释放读锁")

				rw.RUnlock()
				// 当加入读锁时，只有等所有的读锁都解锁完毕才能添加写锁。

				wg.Done()
			}()
		}
		wg.Done()
	}()

	// 写协程:
	go func(data *int) {
		for i := 0; i < 3; i++ {
			go func(data *int) {
				// 上写锁:
				rw.Lock()
				fmt.Println("取得写锁")

				// 模拟写入耗时
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))

				// 写入数据
				*data = rand.Intn(100)
				fmt.Println("写入数据:", *data)
				// 解写锁

				fmt.Println("释放写锁")
				rw.Unlock()
				wg.Done()
			}(data)
		}
		wg.Done()
	}(&data)

	// 等待协程执行完毕
	wg.Wait()
	fmt.Println("结果:", data)
}

// 3 条件变量:
// 条件变量是在mutex的保护下进行的控制协程同步的机制。

func CondTest() {
	var rw sync.RWMutex // 读写锁

	cond := sync.NewCond(rw.RLocker()) // 条件变量(传入读锁)

	var wg sync.WaitGroup
	wg.Add(12)

	// 互斥量:
	data := 0

	// 读多写少:控制条件为在读取时数据大于3

	// 读协程:
	go func() {
		for i := 0; i < 7; i++ {
			go func() {
				rw.RLock()
				fmt.Println("取得读锁")

				// 控制条件:
				for data < 3 {
					// 当控制条件不满足，则在wait处阻塞协程,并且释放传入cond的锁(此处传入读锁，在协程阻塞时自动释放读锁)

					// 使用for循环可以在协程被唤醒时重新判断是否满足条件
					cond.Wait()
				}

				// 读取数据
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100))) // 随机耗时
				fmt.Println("读取数据:", data)

				rw.RUnlock()
				wg.Done()
			}()
		}
		wg.Done()
	}()

	// 写协程

	go func(data *int) {
		for i := 0; i < 3; i++ {
			go func(data *int) {
				rw.Lock()
				fmt.Println("取得写锁")

				// 随机写入耗时
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))

				// 写入数据
				*data += 1
				fmt.Println("写入数据:", *data)

				// 释放
				fmt.Println("释放写锁")
				rw.Unlock()

				cond.Broadcast() // 唤醒所有阻塞的协程:使用for在条件处循环判断，当唤醒后若条件不满足，则继续阻塞

				wg.Done()
			}(data)
		}
		wg.Done()
	}(&data)

	wg.Wait()
	fmt.Println("结果:", data)
}
