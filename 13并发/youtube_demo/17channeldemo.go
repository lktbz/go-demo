package main

import (
	"fmt"
	"sync"
)

//定义等待
var chg8 = sync.WaitGroup{}
//通道读取，循环一直读，当存在数据时，进行读取数据
func main() {
	//定义传递int 类型的channel
	ch := make(chan int, 50) //指定channel容量
	chg8.Add(2)
	//接受数据的
	go func(ch <-chan int) {
			for{
				//判断通道是否被打开
				if i,ok:=<-ch;ok{
					fmt.Println(i)
				}else {
					break
				}
			}
		chg8.Done()
	}(ch)
	//存放数据
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)//关闭通道传输数据，这样就不会产生死锁，不关闭存放数据，获取数据打开通道，就导致了，死锁的产生
		chg8.Done()
	}(ch)

	chg8.Wait()
}
