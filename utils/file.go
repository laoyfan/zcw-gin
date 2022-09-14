package utils

import (
	"io/ioutil"
)

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
