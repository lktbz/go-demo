package main

import (
	"fmt"
	"os"
)

func main() {
//创建文件方式
	create, err := os.Create("test.txt")
	if err!=nil{
		 fmt.Println(err)
		return
	}
	fmt.Println(create)
	create.Close()
	/**
	使用Create()创建文件时：

	如果文件不存在，则创建文件。
	如果文件存在，则清空文件内内容。
	Create创建的文件任何人都可以读写。
	 */
}
