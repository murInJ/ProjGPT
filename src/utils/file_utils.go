package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// IsTextFile checks if the file content appears to be text.
func IsTextFile(path string) bool {
	// Check if the path is a directory
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	if fileInfo.IsDir() {
		return false
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	n, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading file:", err)
		return false
	}

	// Trim the buffer in case the file is shorter than 1024 bytes.
	buffer = buffer[:n]

	// A simple heuristic: if more than 50% of the buffer is printable or space, it's likely a text file.
	printable := 0
	for _, b := range buffer {
		if b == ' ' || b == '\t' || b == '\n' || (b >= ' ' && b <= '~') {
			printable++
		}
	}

	// More than half of the buffer should be printable or space for it to be considered text.
	return float64(printable)/float64(n) > 0.5
}

func ReadFileToString(filePath string) (string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close() // 确保在函数结束时关闭文件

	// 创建 Scanner 对象，用于读取文件
	scanner := bufio.NewScanner(file)

	var content strings.Builder // 使用 strings.Builder 来构建字符串

	// 逐行读取文件内容
	for scanner.Scan() {
		content.WriteString(scanner.Text() + "\n")
	}

	// 检查是否有读取错误
	if err := scanner.Err(); err != nil {
		return "", err
	}

	// 返回文件内容字符串
	return content.String(), nil
}
