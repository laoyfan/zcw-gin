package utils

import (
	"io/ioutil"
)

func GetPathFileNames(path string) (names []string) {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if !file.IsDir() {
			names = append(names, file.Name())
		}
	}
	return
}
