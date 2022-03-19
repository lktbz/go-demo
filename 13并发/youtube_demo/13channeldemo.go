package main

import (
	"fmt"
	"sync"
)
//定义等待
var chg4=sync.WaitGroup{}

func main() {
	//定义传递int 类型的channel
	ch:=make(chan int)
	chg4.Add(2)
	//接受数据的
	go func() {
		i:=<-ch  //从channel 中获取数据
		fmt.Println("go1--",i)
		ch<-27
		chg4.Done()
	}()
		//存放数据
		go func() {
			ch<-42 //将刷剧放进channel  将数据刷入通道导致死锁发生
		fmt.Println("go2--",<-ch)
			chg4.Done()
		}()

	chg4.Wait()
}
