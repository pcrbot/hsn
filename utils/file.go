package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func IsExist(fileAddr string) bool {
	// 读取文件信息，判断文件是否存在
	_, err := os.Stat(fileAddr)
	if err != nil {
		if os.IsExist(err) { // 根据错误类型进行判断
			return true
		}
		return false
	}
	return true
}

// 判断目录是否存在
func IsDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func GetFiles(folder string) []string {
	var r []string
	files, _ := ioutil.ReadDir(folder) //specify the current dir
	for _, file := range files {
		if file.Name() == "__pycache__" { // ignore pycache
			continue
		}

		if file.IsDir() {
			r = append(r, GetFiles(folder+"/"+file.Name())...)
		} else {
			r = append(r, fmt.Sprint(folder+"/"+file.Name()))
		}
	}
	return r
}
