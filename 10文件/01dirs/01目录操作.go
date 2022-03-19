package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//isExistsDir()
	//isExistsDirTwo()
	isExistsDirAll()
}

/**
目录不存在则创建目录
*/
func isExistsDir() {
	_, err := os.ReadDir("e:/go")
	if err != nil {
		os.Mkdir("e:/go", os.ModeDir)

	}
	dir, _ := os.ReadDir("e:/go")
	for _, entry := range dir {
		fmt.Println(entry.Name())
	}
}

/**
  健壮性修改
*/
func isExistsDirTwo() {
	_, err := os.ReadDir("e:/go")
	if err != nil {
		os.Mkdir("e:/go", os.ModeDir)

	}
	dir, _ := os.ReadDir("e:/go")
	//不存在的目录不进行遍历操作
	for i := 0; len(dir) > 0 && i < len(dir); i++ {
		fmt.Println(dir[i].Name())
	}
}
//创建多层文件
func isExistsDirAll() {
	_, err := os.ReadDir("e:/gos")
	if err != nil {
		os.MkdirAll("e:/gos/ggo/g/o/gg", os.ModeDir)
	}
	dir, _ := os.ReadDir("e:/gos")
	//不存在的目录不进行遍历操作
	for i := 0; len(dir) > 0 && i < len(dir); i++ {
		fmt.Println(dir[i].Name())
	}
}
//读取目录
func readDir() {
	dir, err := os.ReadDir("d:/go")
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range dir {
		fmt.Println(entry.Name())
	}
}
func ioutils()  {
		dir, _ := ioutil.ReadDir("e:/go")
	for _, info := range dir {
		os.RemoveAll(info.Name())
	}
}

//百度出来的方法
//判断文件文件夹是否存在
func isFileExist(path string) (bool, error) {
	fileInfo, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}
	//我这里判断了如果是0也算不存在
	if fileInfo.Size() == 0 {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}