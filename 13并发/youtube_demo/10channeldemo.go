package main

import (
	"fmt"
	"sync"
)
//定义等待
var chg=sync.WaitGroup{}
func main() {
	//定义传递int 类型的channel
	ch:=make(chan int)
	chg.Add(2)
	//接受数据的
	go func() {
		i:=<-ch  //从channel 中获取数据
		fmt.Println(i)
		chg.Done()
	}()
	//存放数据
	go func() {
		i:=42
		ch<-i   //将刷剧放进channel
		i=27 //修改i的值，查看返回数据
		chg.Done()
	}()
	chg.Wait()
}
