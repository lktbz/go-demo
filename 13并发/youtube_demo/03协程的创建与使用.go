package main

import (
	"fmt"
	"time"
)

func main() {
  var msg="hello"
	go func() {
		fmt.Println(msg)
	}()
  msg="Goodbye" //变量被修改，导致，直接打印goodbye ，跟期望的又差别
  time.Sleep(100*time.Millisecond)
}
