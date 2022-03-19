package main

import (
	"fmt"
	"sync"
)
//定义等待
var chg5=sync.WaitGroup{}

func main() {
	//定义传递int 类型的channel
	ch:=make(chan int)
	chg5.Add(2)
	//接受数据的
	go func(ch <-chan int) {
		i:=<-ch  //从channel 中获取数据
		fmt.Println("go1--",i)
		//ch<-27
		chg5.Done()
	}(ch)
		//存放数据
		go func(ch chan<-int) {
			ch<-42 //将刷剧放进channel  将数据刷入通道导致死锁发生
		//fmt.Println("go2--",<-ch)
			chg5.Done()
		}(ch)

	chg5.Wait()
}
