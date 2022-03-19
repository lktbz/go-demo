package main

import (
	"fmt"
)

/**
   channel 关闭
 */


func main() {
	ch:=make(chan int ,20)
     go func(ch chan int) {
		 for i := 0; i < 10; i++ {
			 ch<-i
		 }
		 close(ch)// 关闭通道
	 }(ch)

	//空循环判断通道是否关闭
	for{
		if  num,ok:= <-ch;ok==true{
			fmt.Println("读到的数据为,",num)
		}else {
			fmt.Println("通道关闭")
			break
		}
	}
}
