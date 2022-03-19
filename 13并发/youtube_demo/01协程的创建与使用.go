package main

import (
	"fmt"
	"time"
)
//函数
func  sayHello()  {
	fmt.Println("hello")
}
func main() {
  go sayHello() //使用协程
  //主线程睡眠，不然没有执行效果
  time.Sleep(5*time.Second) //睡眠5秒
}
