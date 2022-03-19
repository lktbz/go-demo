package main

import (
	"fmt"
	"time"
)

/**
  现在有100个并发，限制在5个
 */

func timeMore(ch chan string){
	//执行前先注册，写不进去就会阻塞
	ch<-"任务"
	 fmt.Println("模拟耗时操作")
	  time.Sleep(time.Second)
	//任务执行完毕，则管道中销毁一个任务
	<-ch
}
func main() {
		ch:=make(chan string,5)
	for i := 0; i <100 ; i++ {
		go timeMore(ch)
	}
	for  {
		time.Sleep(time.Second*5)
	}
}
