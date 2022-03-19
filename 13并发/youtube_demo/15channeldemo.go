package main

import (
	"fmt"
	"sync"
)
//定义等待
var chg6=sync.WaitGroup{}

func main() {
	//定义传递int 类型的channel
	ch:=make(chan int,50)  //指定channel容量
	chg6.Add(2)
	//接受数据的
	go func(ch <-chan int) {
		i:=<-ch  //从channel 中获取数据
		fmt.Println("go1--",i)
		//i=<-ch
		//fmt.Println(i)  到现在都不是正确的处理channel 方式


		chg6.Done()
	}(ch)
		//存放数据
	go func(ch chan<-int) {
			ch<-42
		    ch<-27
			chg6.Done()
		}(ch)

	chg6.Wait()
}
