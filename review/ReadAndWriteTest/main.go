// 编写CLI程序，用户输入路径，读取路径下所有目录及其文件，并打开指定类型的文件，实现对文件的读写操作。
// 还需实现目录和文件的创建删除操作。

package main

import (
	"fmt"
	// "io"      // 用于处理输入输出
	"os" // 用于处理文件和目录

	"review/ReadAndWriteTest/tools"
	"review/ReadAndWriteTest/tools/stringResolve"
)

var (
	DirPath  string         // 储存路径
	DirEntry []os.DirEntry  // os.DirEntry接口实现4个方法：Name() string，IsDir() bool，Type() os.FileMode，Info() (os.FileInfo, error) 此处得到的是包含os.DirEntry的切片
	DirInfo  map[int]string // 储存目录名

	err            error
	filePath       string // 储存文件路径
	activeFilePath string // 储存当前选中的文件路径
	activeDirPath  string // 储存当前选中的目录路径
)

func main() {
	// 等待用户输入
	fmt.Println("输入文件路径")

	// 标准输入：循环监听用户输入，直到路径合法
	for {
		fmt.Scanln(&DirPath) // 使用指针，函数访问指针并储存到变量

		// 读取路径下目录及其文件
		DirInfo, err, DirEntry = tools.ReadDir(DirPath)
		if err != nil {
			fmt.Println("读目录失败，清重新输入：", err)
			continue // 使用continue直接开始新的循环
		} else {
			fmt.Println("读取到路径:")
			break
		}
	}
	// 等待操作

	for {
		var pt string
		tools_stringResolve.RangeMap(DirInfo)
		fmt.Println("请输入: 1.读取文件 2.写入文件 3.创建文件 4.删除文件 5.创建目录 6.进入目录 7.删除目录 8.退出程序")
		fmt.Scanln(&pt)
		switch pt {
		case "1":
			var (
				c string
			)
			acFilePathChange()
			c, err = tools.ReadFile(activeFilePath)
			if err != nil {
				fmt.Println("读取文件失败：", err)
			} else {
				fmt.Println(c)
			}
		case "2":
			var (
				c string
			)
			acFilePathChange()
			fmt.Println("请输入内容:")
			fmt.Scanln(&c)
			err = tools.WriteFile(activeFilePath, c, false)
			if err != nil {
				fmt.Println("写入文件失败：", err)
			}
		case "3":
			var (
				c  string
				fn string
			)
			fmt.Println("指定目录")
			acDirPathChange()

			fmt.Println("请输入文件名:")
			fmt.Scanln(&fn)
			activeFilePath = activeDirPath + "/" + fn + ".txt"

			fmt.Println("请输入文件内容:")
			fmt.Scanln(&c)
			err = tools.WriteFile(activeFilePath, c, true)
			if err != nil {
				fmt.Println("创建文件失败：", err)
			}

		case "4":
			acFilePathChange()
			err = os.Remove(activeFilePath)
			if err != nil {
				fmt.Println("删除文件失败：", err)
			}
		case "5":

		default:
			fmt.Println("请输入正确选项")
			fmt.Println(pt)
		}
	}
}

//############################### 辅助函数 ##############

func acFilePathChange() {
	for {
		var (
			n int
		)
		fmt.Println("请输入文件序号:")
		fmt.Scanln(&n)
		if DirInfo[n+len(DirInfo)/2] == "Y" {
			fmt.Println("无法选取目录，请重新选择")
		} else if DirInfo[n+len(DirInfo)/2] == "N" {
			activeFilePath = DirPath + "/" + DirInfo[n]
			return
		} else {
			fmt.Println("文件不存在，请重新选择")
		}
	}
}

func acDirPathChange() {
	for {
		var (
			n int
		)
		fmt.Println("请输入目录序号:")
		fmt.Scanln(&n)
		if DirInfo[n+len(DirInfo)/2] == "Y" {
			activeDirPath = DirPath + "/" + DirInfo[n]
			return
		} else if DirInfo[n+len(DirInfo)/2] == "N" {
			fmt.Println("无法选取文件，请重新选择")
		} else {
			fmt.Println("目录不存在，请重新选择")
		}
	}
}
