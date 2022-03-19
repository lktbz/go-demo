package main

import (
	"fmt"
	"sync"
)
//定义等待
var chg3=sync.WaitGroup{}
//产生了死锁，通道的容量是1，没有设置，
func main() {
	//定义传递int 类型的channel
	ch:=make(chan int)
	//接受数据的
	go func() {
		i:=<-ch  //从channel 中获取数据
		fmt.Println(i)
		chg3.Done()
	}()
	//5个存放，一个接受
	for i := 0; i < 5; i++ {
		chg3.Add(2)
		//存放数据
		go func() {
			ch<-42 //将刷剧放进channel  将数据刷入通道导致死锁发生
			chg3.Done()
		}()
	}
	chg3.Wait()
}
