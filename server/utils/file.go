package utils

import (
	"io"
	"os"
)

// CopyFile 复制文件函数
func CopyFile(src, dst string) error {
	// 打开源文件
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// 创建目标文件
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 使用io.Copy复制内容
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// 确保写入目标文件的内容
	return destFile.Sync()
}
