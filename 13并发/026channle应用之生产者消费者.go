package main

import (
	"fmt"
	"time"
)
//生产者
func  producer(ch chan<-int)()  {
	i:=1
	for{
		ch<-i
		fmt.Println("send:",i)
		i++;
		time.Sleep(time.Second*1)
	}
}
//消费数据
func consumer(ch<-chan int){
	for  {
		v:=<-ch
		fmt.Println("Receive:",v)
		time.Sleep(time.Second*2)
	}
}
func main() {
  ch:=make(chan int ,5)
  go producer(ch)
  go consumer(ch)
  for{}
}
