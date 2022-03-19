package main

import (
	"fmt"
	"time"
)

/**
Go语言中的 select 关键字，可以同时响应多个通道的操作，在多个管道中随机选择一个能走通的路！
语法：
select {
	case 操作1:
		响应操作1
	case 操作2:
		响应操作2
	...
	default:
		没有操作的情况
}
*/
//如果有这样的需求，两个管道中只要有一个管道能够取出数据，那么就使用该数据：
func fn1(ch chan string){
	time.Sleep(time.Second*3)
	ch<-"fn1111"
}
func fn2(ch chan string)  {
	time.Sleep(time.Second*1)
	ch<-"fn2222"
}
func main() {
     ch1:=make(chan string)
     go fn1(ch1)
     ch2:=make(chan string)
     go fn2(ch2)
	select {
	 case r1:=<-ch1:
		fmt.Println("r1=:",r1)
	 case r2:=<-ch2:
	 	fmt.Println("r2=:",r2)
	}
	/**
	select支持default，如果select没有一条语句可以执行，即所有的通道都被阻塞，那么有两种可能的情况：
	- 如果给出了default语句，执行default语句，同时程序性的执行会从select语句后的语句中恢复
	- 如果没有default语句，那么select语句将被阻塞，直到至少有一个通信可以进行下去
	- 所以一般不写default语句

	当然，在一些场景中（for循环使用select获取channel数据），如果channel被写满，也可能会执行default。

	注意：select中的case必须是I/O操作。
	*/
}
