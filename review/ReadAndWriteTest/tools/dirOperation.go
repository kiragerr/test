package tools

import (
	// "errors"
	// "fmt"
	// "io"
	"os"
)

func ReadDir(dirPath string) (map[int]string, error, []os.DirEntry) {
	RString := map[int]string{}
	DirEntry, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err, nil
	}
	// 遍历路径下所有目录
	for i, entry := range DirEntry {
		if entry.IsDir() {
			RString[i] = "/" + entry.Name()
			RString[i+len(DirEntry)] = "Y"
		} else {
			RString[i] = entry.Name()
			RString[i+len(DirEntry)] = "N"
		}
	}
	// fmt.Printf("%v", DirEntry)
	return RString, nil, DirEntry
}
