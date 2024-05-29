package parser

import (
	"ProjGPT/src/utils"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getFileContentString(root string) (string, error) {
	var b bytes.Buffer // 使用 bytes.Buffer 来累积输出
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// 指定要跳过的子路径
		skipPathList := []string{".git", ".idea", ".vscode"}
		for _, skipPath := range skipPathList {
			// 检查当前路径是否是指定子路径的子目录
			if strings.HasPrefix(path, filepath.Join(root, skipPath)) {
				// 如果是，返回 filepath.SkipDir 以跳过该目录下的所有文件
				return nil
			}
		}

		if utils.IsTextFile(path) {
			fileString, err := utils.ReadFileToString(path)
			if err != nil {
				return err
			}
			b.WriteString(strings.Repeat("$", 20))
			b.WriteString("\n")
			b.WriteString(path)
			b.WriteString("\n")
			b.WriteString(strings.Repeat("-", 20))
			b.WriteString("\n")
			b.WriteString(fileString)
			b.WriteString("\n")
			b.WriteString(strings.Repeat("$", 20))
			b.WriteString("\n")
		}
		return nil
	})

	if err != nil {
		return "", errors.New(fmt.Sprintf("Error walking the path: %v\n", err))
	}
	return b.String(), nil // 返回累积的字符串
}
