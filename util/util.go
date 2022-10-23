package util

import (
	"errors"
	"io/ioutil"
	"os"
)

// PathExists 判断路径存在
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetPathFileNames 获取文件夹下文件名称
func GetPathFileNames(path string) (names []string) {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if !file.IsDir() {
			names = append(names, file.Name())
		}
	}
	return
}
