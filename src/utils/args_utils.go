package utils

import (
	"fmt"
	"os"
)

func Check_root(rootPtr *string) {
	fileInfo, err := os.Stat(*rootPtr)
	// 检查路径是否存在
	if os.IsNotExist(err) {
		fmt.Println("Error: specified root path does not exist.")
		os.Exit(1)
		return
	}
	// 检查路径是否为目录
	if err != nil {
		fmt.Println("Error: failed to stat the root path:", err)
		os.Exit(1)
		return
	}
	if !fileInfo.IsDir() {
		fmt.Println("Error: the specified root path is not a directory:", *rootPtr)
		os.Exit(1)
		return
	}
}
