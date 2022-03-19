package main

import (
	"fmt"
	"os"
)

/**
打开文件有两种方式：

Open()：以只读的方式打开文件，若文件不存在则会打开失败
OpenFile()：打开文件时，可以传入打开方式，该函数的三个参数：
参数1：要打开的文件路径
参数2：文件打开模式，如 O_RDONLY，O_WRONGLY，O_RDWR，还可以通过管道符来指定文件不存在时创建文件
参数3：文件创建时候的权限级别，在0-7之间，常用参数为6


常用的文件打开模式：

	O_RDONLY 	int = syscall.O_RDONLY		// 只读
	O_WRONGLY	int = syscall.O_WRONGLY		// 只写
	O_RDWR 		int = syscall.O_RDWR		// 读写
	O_APPEND 	int = syscall.O_APPEND		// 写操作时将数据追加到文件末尾
	O_CREATE 	int = syscall.O_CREATE		// 如果不存在则创建一个新文件
	O_EXCL 		int = syscall.O_EXCL		// 打开文件用于同步I/O
	O_TRUNC		int = syscall.O_TRUNC		// 如果可能，打开时清空文件
 */
func main() {
  //openFileOne()
	openFileTwo()
}
func openFileTwo()  {
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_RDWR, os.ModeAppend)
    if err!=nil{
		fmt.Println("open file err:",err)
	}else
	{
		fmt.Println(file)
	}
	file.Close()

}
func openFileOne()  {
	open, err := os.Open("test.txt")
	if err ==nil{
		fmt.Println("文件不存存在")
		return
	}else {
		fmt.Println("打开文件成功",open)
	}
	open.Close()
}