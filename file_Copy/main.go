package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// RDandWR("./data/README.txt", "./data/README_copy.txt")
	// OSFileReadFrom("./data/README.txt", "./data/README_copy.txt")
	IOCopyFile("./data/README.txt", "./data/README_copy.txt")
}

// 分别使用
// 1 读取+写入
// 2 os.File.ReadFrom
// 3 io.Copy
// 复制文件。

// 1 读取 + 写入

func RDandWR(name1 string, name2 string) {
	file1, err := os.OpenFile(name1, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("file1 opened")
		defer file1.Close()
		file2, err := os.OpenFile(name2, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("file2 opened")
			buffer, err := io.ReadAll(file1)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Writing to file2")
				n, err := file2.Write(buffer)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(n, "bytes written to file2")
					file2.Close()
				}
			}
		}
	}
}

// 2 os.File.ReadFrom

func OSFileReadFrom(name1 string, name2 string) {
	file1, err := os.OpenFile(name1, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("file1 opened")
		defer file1.Close()
		file2, err := os.OpenFile(name2, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("file2 opened")
			n, err := file2.ReadFrom(file1)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(n, "bytes copied to file2")
				file2.Close()
			}
		}
	}
}

// 3 io.Copy

func IOCopyFile(name1 string, name2 string) {
	file1, err := os.OpenFile(name1, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("file1 opened")
		defer file1.Close()
		file2, err := os.OpenFile(name2, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("file2 opened")
			n, err := io.Copy(file2, file1)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(n, "bytes copied to file2")
				file2.Close()
			}
		}
	}
}
