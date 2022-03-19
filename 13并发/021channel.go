package main

import (
	"fmt"
	"time"
)

/**
   无缓冲的channel
无缓冲的channel是阻塞读写的，必须写端与读端同时存在，写入一个数据，则能读出一个数据：
 */

func  writ(ch chan int)  {
     ch<-100
	fmt.Printf("ch addr：%v\n", ch) // 输出内存地址
	ch<-200
	fmt.Printf("ch addr：%v\n", ch) // 输出内存地址
	ch<-300  // 该处数据未读取，后续操作直接阻塞
	fmt.Printf("ch addr：%v\n", ch) // 没有输出
}
func rea(ch chan int){
	// 只读取两个数据
	fmt.Printf("取出的数据data1：%v\n",<-ch)
	fmt.Printf("取出的数据data2：%v\n",<-ch)
}
func main() {
	var ch chan int // 声明一个无缓冲的channel
	ch=make(chan int)
	// 向协程中写入数据
    go writ(ch)
	// 向协程中读取数据
	go rea(ch)
	// 防止主go程提前退出，导致其他协程未完成任务
	time.Sleep(time.Second * 3)
	/**
	注意：无缓冲通道的收发操作必须在不同的两个goroutine间进行，
	因为通道的数据在没有接收方处理时，数据发送方会持续阻塞，所以通道的接收必定在另外一个 goroutine 中进行。
	 */
}
