package main

import (
	"fmt"
)

func main() {
  f:=func(){
  	 msg:="hello go"
  	fmt.Println(msg)
  }
  //调用f 函数
  f()
  var fs func()= func() {
	  fmt.Println("继续变种调用")
  }
  fs()
}
