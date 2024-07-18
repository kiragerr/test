package context_learn

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 使用timerCtx模拟web请求的超时处理
// timer在timerCtx创建之初就开始及时。

func TimerCtxTest() {
	// 包装根cancelCtx节点
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)

	// 开启两个web请求。

	// 两个请求允许超时时间不同
	go HttpHandler1(ctx, time.Second*2, &wg)
	go HttpHandler2(ctx, time.Second*1, &wg)

	// 设置主协程等待时间
	time.Sleep(3 * time.Second)
	cancel() // 关闭所有请求
	wg.Wait()
	fmt.Println("All requests are done.")

}

// web请求1：
func HttpHandler1(ctx context.Context, timerout time.Duration, wg *sync.WaitGroup) {
	// 包装timerCtx节点
	timerCtx, cancel := context.WithTimeout(ctx, timerout)
	defer cancel()
	// 模拟处理web请求
WebLoopL:
	for {
		select {
		case <-timerCtx.Done():
			fmt.Println("请求停止:", timerCtx.Err())
			break WebLoopL
		default:
			fmt.Println("处理web请求中...")
			time.Sleep(time.Millisecond * 500)
		}
	}
	fmt.Println("web1请求结束")
	wg.Done()
}

// web请求2：

func HttpHandler2(ctx context.Context, timeout time.Duration, wg *sync.WaitGroup) {
	timerCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel() // 在流程结束后手动取消上下文

	// 模拟web超时

	for {
		select {
		case <-timerCtx.Done():
			fmt.Println("请求停止:", timerCtx.Err())
			wg.Done()
			fmt.Println("web2请求结束")
			return
		}
	}

}
