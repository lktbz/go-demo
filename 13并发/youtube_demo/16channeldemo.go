package main

import (
	"fmt"
	"sync"
)

//定义等待
var chg7 = sync.WaitGroup{}

func main() {
	//定义传递int 类型的channel
	ch := make(chan int, 50) //指定channel容量
	chg7.Add(2)
	//接受数据的
	go func(ch <-chan int) {
		//使用循环的方式接受channel 的数据
		for i := range ch {
			fmt.Println(i)
		}
		chg7.Done()
	}(ch)
	//存放数据
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)//关闭通道传输数据，这样就不会产生死锁，不关闭存放数据，获取数据打开通道，就导致了，死锁的产生
		chg7.Done()
	}(ch)

	chg7.Wait()
}
