package context_learn

import (
	"context"
	"fmt"
	"sync"
)

// 练习：使用ValueCtx传递上下文信息

type key1 struct{}
type key2 struct{} // 类型声明可以在函数外部(全局域)。

func ValueCtxTest() {
	// 创建empty context
	ctx := context.Background()

	// 创建WithValue的ValueCtx
	// ValueCtx封装两个子节点
	valuectx1 := context.WithValue(ctx, key1{}, "value1")
	valuectx2 := context.WithValue(ctx, key2{}, "value2")

	// 创建WaitGroup
	wg := sync.WaitGroup{}
	wg.Add(3)

	// 开启两个协程，其中一个为嵌套

	go goRoutine1(valuectx1, &wg)
	go goRoutine2(valuectx2, &wg)

	// 等待协程完毕
	wg.Wait()
	fmt.Println("All goroutines done")
}

func goRoutine1(ctx context.Context, wg *sync.WaitGroup) {
	// 从ValueCtx中获取key1的值
	fmt.Println("goRoutine1: key1 value is", ctx.Value(key1{}))
	wg.Done()
}

func goRoutine2(ctx context.Context, wg *sync.WaitGroup) {
	// 开启嵌套协程
	go func(ctx context.Context) {
		// 从透传的ValueCtx获取key2值
		fmt.Println("goRoutine2: key2 value is", ctx.Value(key2{}))
		wg.Done()
	}(ctx)
	wg.Done()
}
