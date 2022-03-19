package main

import (
	"fmt"
	"time"
)

/**
	通道数据的遍历
	channel只支持 for--range 的方式进行遍历：
*/
func main() {
	ch:=make(chan int)
   go func() {
	   for i := 0; i <=3 ; i++ {
		   ch<-i
		   fmt.Println("i的值为：",i)
		   time.Sleep(time.Second)
	   }
   }()

	for  data := range ch {
		fmt.Println("data==",data)
		if data==3 {
			break
		}
	}
}
