package main

import (
	"fmt"
	"time"
)

/**
  协程的开启
 */
func  say(s string)  {
	for i := 0; i <3 ; i++ {
		fmt.Println(s)
	}
}
func world(s string)  {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}
func main() {
   go say("go")   //开启协程
   go world("lala")
   say("GO")   //以普通方式允许线程
   time.Sleep(5*time.Second) //睡眠5秒
}
