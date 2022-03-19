package main

import (
	"fmt"
	"time"
)
/**
  对03进行修改
 */
func main() {
  var msg="hello"
	go func(msg string) {
		fmt.Println(msg)  //打印的结果为hello
	}(msg)
  msg="Goodbye"
  time.Sleep(100*time.Millisecond)
}
