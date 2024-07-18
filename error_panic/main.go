package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	WrapExample()
	IsExample()
	AsExample()
}

// 创建一个链式错误，并使用Wrap解包,使用Is和As方法进行错误类型判断和类型转换

// 1 使用 Wrap 解包
func WrapExample() {
	// 创建链式错误
	err := errors.New("error1")
	WrapErr := fmt.Errorf("error2: %w", err)

	// 使用Wrap解包
	fmt.Printf("WrapErr: %v\n", errors.Unwrap(WrapErr)) // WrapErr: error1
}

// 2 使用 Is 判断是否包含某种错误
func IsExample() {
	// 创建链式错误
	err := errors.New("error1")
	WrapErr1 := fmt.Errorf("error2: %w", err)
	WrapErr2 := fmt.Errorf("error3: %w", WrapErr1)

	if errors.Is(WrapErr2, WrapErr1) {
		fmt.Println("WrapErr2 is WrapErr1")
	} else {

		fmt.Println("WrapErr2 is not WrapErr1")
	} // WrapErr2 is WrapErr1
}

// 3 使用 As 转换错误类型

func AsExample() {
	// 实例化链式错误:
	// 1 创建一个自定义错误
	err := Wrap1("error1")
	// 2 创建错误结构体接口(target)
	var targetErr *MyError
	// 3 使用As方法转换错误类型
	// target必须传入指针类型，用于之后的引用。
	// 当识别为子类型时，将第一个匹配的错误链的值复制到target指向的结构体中。
	if errors.As(err, &targetErr) {
		fmt.Printf("targetErr: %#v\n", targetErr) //格式化输出完整targetErr结构体
		fmt.Printf("Time: %v\n", targetErr.Time)  // 输出MyError结构体中的Time字段
	} else {
		fmt.Println("targetErr is not *MyError")
	}
}

// 错误结构:
type MyError struct {
	Msg  string
	Time time.Time // Time类型
}

// error 接口类型要求内部包含Error() 方法。
func (e *MyError) Error() string {
	return e.Msg
}

// 创建自定义错误函数
func NewMyError(msg string) error {
	return &MyError{
		Msg:  msg,
		Time: time.Now(),
	}
	// 一般这类返回一个结构体或其他结构的函数，叫做工厂函数。
	// go语言中，一般使用NewXXX函数作为工厂函数的命名格式。
}

// 创建链式错误
func Wrap1(msg string) error {
	return fmt.Errorf("error2: %w", Wrap2(msg))
}

func Wrap2(msg string) error {
	return NewMyError(msg)
}
