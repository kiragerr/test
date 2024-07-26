package tools

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// 文件读取功能
func ReadFile(fp string) (string, error) {

	var (
		fileContent string
	)

	// 打开文件
	f, err := os.Open(fp)
	defer f.Close() // 注意关闭文件解除占用

	if err != nil {
		return "", err
	} else {

		// 读取文件内容
		buffer, err := func(f *os.File) ([]byte, error) {
			b := make([]byte, 1024) // 分配初始缓冲区大小

			// 循环读取文件内容(每一次读取一部分(读到缓冲区满开始下一次读取))
			for {
				// 判断缓冲区是否溢出，当缓冲区内容=容量时，使用切片特性扩容
				if len(b) == cap(b) {
					b = append(b, 0)[:len(b)]
				}

				n, err := f.Read(b[len(b):cap(b)]) // os.Read方法会调用切片并直接写入到切片中，返回写入的字节数和err
				// 此处切片要传入切片后半空余的部分让Read写入。

				b = b[:len(b)+n] // 拼接写入的内容

				// 判断是否读取完毕/出错
				if err != nil {
					// 使用errors.Is判断错误是否为io.EOF类型
					if errors.Is(err, io.EOF) {
						err = nil
						break
					}
				}
			}
			return b, err
		}(f)

		// 将buffer转换为字符串

		if err != nil {
			return "", err
		} else {
			fileContent = string(buffer)
			fmt.Println("文件读取完毕")
			return fileContent, nil
		}
	}
}

// 写到文件

func WriteFile(fp string, content string, IsCreate bool) error {
	// 根据是否需要创建新文件，决定文件的打开方式

	// 文件打开模式
	var (
		RWMode     int = os.O_RDWR | os.O_APPEND               // 普通读写模式
		CreateMode int = os.O_RDWR | os.O_APPEND | os.O_CREATE // 创建文件模式
	)
	if IsCreate {
		RWMode = CreateMode
	}

	// 打开文件
	f, err := os.OpenFile(fp, RWMode, 0666)
	defer f.Close()
	if err != nil {
		return err
	} else {
		// 写入文件
		n, err := io.WriteString(f, content) // 传入io.Writer实现
		if err != nil {
			return err
		} else {
			fmt.Println("写入完毕,写入字节:", n)
			return nil
		}
	}
}
