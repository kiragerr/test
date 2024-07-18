package main

import (
	"fmt"
	// chTest "multithreading/channel_RangeAndSelect"
	// ctxTest "multithreading/context_Learn"
	// "multithreading/mutex"
	// syncTest "multithreading/sync"
	mutexTest "multithreading/mutex"
)

func main() {

	// // p1 管道
	// mutex.NonBufferedChannelExample()
	// fmt.Println("-------------")
	// mutex.BufferedChannelExample()

	// // 输出结果: s:0 1 2 r:0 s:3 r:1 ....
	// // 由于读取耗时1s比写入慢，所以在管道有缓冲时先执行写入，当缓冲区阻塞，等待读取，读取完毕时，在下一个读取时间之前又可以继续写入(先执行快的操作，提升效率)

	// // p2 Range/Select channel
	// fmt.Println("--------------")
	// chTest.RangeChannel()
	// fmt.Println("--------------")
	// chTest.SelectExample()

	// // p3 sync包
	// fmt.Println("---------------")
	// syncTest.WaitGroupExample()

	// // p4 context上下文
	// fmt.Println("---------------")
	// ctxTest.ValueCtxTest()
	// fmt.Println("---------------")
	// ctxTest.TimerCtxTest()

	// p5 mutex
	fmt.Println("---------------")
	mutexTest.RWMutexTest()
	fmt.Println("---------------")
	mutexTest.CondTest()
}
