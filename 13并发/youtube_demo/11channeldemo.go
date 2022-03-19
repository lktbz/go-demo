package main

import (
	"fmt"
	"sync"
)
//定义等待
var chg2=sync.WaitGroup{}
func main() {
	//定义传递int 类型的channel
	ch:=make(chan int)
	//5个接受，5个发送
	for i := 0; i < 5; i++ {
		chg2.Add(2)
		//接受数据的
		go func() {
			i:=<-ch  //从channel 中获取数据
			fmt.Println(i)
			chg2.Done()
		}()
		//存放数据
		go func() {
			ch<-i   //将刷剧放进channel
			chg2.Done()
		}()
	}


	chg2.Wait()
}
