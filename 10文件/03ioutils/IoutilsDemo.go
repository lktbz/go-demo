package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
 //readFile()
	//ReadsFile()
	WriteFiles()
}
func WriteFiles(){
	s:="test.txt"
    data:=[]byte("hello ioutils write")

	err := ioutil.WriteFile(s, data, 0644)
	if err != nil {
		fmt.Println("文件写入失败")
	}
	fmt.Println("文件写入完毕")
}
func ReadsFile()  {
	dir, err := ioutil.ReadDir("D:\\go\\workspace\\go\\03数据类型")
	if err!=nil{
		fmt.Println("文件读取失败。。。。")
	}
	for _, info := range dir {
		fmt.Println("目录为："+info.Name())
	}
}
func readFile()  {
	file, err := ioutil.ReadFile("test.txt")
	if err!=nil {
		fmt.Println("读取文件失败")
	}
	fmt.Println(string(file))
}