package main

import (
	"fmt"
	"os"
)
/**

   服务文件 的函数定义
   func Stat(name string) (FileInfo, error) {
	testlog.Stat(name)
	return statNolog(name)
	}
 */
func main() {
	stat, err := os.Stat("a.txt")
	if err!=nil{
		fmt.Println("出错了",err.Error())
	}else {
		fmt.Println(stat.Name())
	}

}
