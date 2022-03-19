package main

import (
	"fmt"
)

func main() {
  //fmt.Println("sa")
  sayMessage()
  saymessage("lktbz")
	for i := 0; i <5 ; i++ {
		Saymessage("hello go",i)
	}
}
//没有参数的函数
func sayMessage(){
	fmt.Println("调用了say message")
}
//传递一个参数
func saymessage(msg string)  {
	fmt.Println(msg)
}
//两个参数
func Saymessage (msg string,idx int){
  fmt.Println(msg)
  fmt.Println("value is:",idx)
}