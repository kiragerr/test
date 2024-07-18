package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// 分别使用三种不同的文件读取方式读取文件。
// 1. os.File Read()
// 2. os.ReadFile()
// 3. io.ReadAll()

func main() {
	UseOsFileReader("README.txt")
	UseOsReadFile("README.txt")
	UseIoReadAll("README.txt")
}

// 1. 使用os.File文件类型下Read方法读取文件。

func IoReadFile(f *os.File) ([]byte, error) {
	// 1 创建buffer切片
	b := make([]byte, 0, 512)
	// b := make([]byte, 0, 2) // 测试扩容

	// 2 循环判断切片容量是否已满
	for {
		if len(b) == cap(b) {
			// 满则扩容
			b = append(b, 0)[:len(b)]
			fmt.Println("扩容")
		}

		// 3 读取文件
		offset, err := f.Read(b[len(b):cap(b)]) // offset为读取字节数

		// 4 处理读取后数据
		b = b[:len(b)+offset]

		// 5 判断错误类型

		if err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
			}
			// 返回读取完毕的字节切片和错误信息
			return b, err

		}

	}
}

func UseOsFileReader(ad string) {
	file, err := os.OpenFile(ad, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("文件打开失败:", err)
	}
	buffer, err := IoReadFile(file)

	if err != nil {
		fmt.Println("文件读取失败:", err)
	} else {
		fmt.Println(string(buffer))

		// 关闭文件
		file.Close()
	}

}

// 2 使用os.ReadFile()函数读取文件。
func UseOsReadFile(ad string) {
	buffer, err := os.ReadFile(ad)
	if err != nil {
		fmt.Println("文件读取失败:", err)
	} else {
		fmt.Println(string(buffer))
	}
}

// 3 使用io.ReadAll()函数读取文件。
func UseIoReadAll(ad string) {
	// 使用os.OpenFile打开文件，并得到os.File接口
	file, err := os.OpenFile(ad, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("文件打开失败:", err)
	} else {
		buffer, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("文件读取失败:", err)
		} else {
			// 直接写出buffer流
			fmt.Println(buffer)
			// 关闭文件
			file.Close()
		}
	}
}
