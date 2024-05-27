package parser

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// getStructTreeString 返回目录树结构的字符串表示形式
func getStructTreeString(root string) (string, error) {
	var b bytes.Buffer // 使用 bytes.Buffer 来累积输出
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 构建当前节点的相对路径
		relPath, _ := filepath.Rel(root, path)
		// 计算缩进层数
		indent := strings.Count(relPath, string(filepath.Separator)) + 1
		// 格式化输出，使用缩进来表示目录结构
		prefix := strings.Repeat("    ", indent-1) // 缩进
		if info.IsDir() {
			fmt.Fprintf(&b, "%s%s/\n", prefix, info.Name())
		} else {
			fmt.Fprintf(&b, "%s%s\n", prefix, info.Name())
		}
		return nil
	})

	if err != nil {
		return "", errors.New(fmt.Sprintf("Error walking the path: %v\n", err))
	}
	return b.String(), nil // 返回累积的字符串
}
