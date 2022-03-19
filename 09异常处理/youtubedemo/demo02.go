package main

import "fmt"

func main() {
fmt.Println("start")
//使用内部处理机制 ，出错误，并终止程序运行
panic("something is bad")
fmt.Println("end")
}
