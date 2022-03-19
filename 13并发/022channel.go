package main

import (
	"fmt"
	"time"
)

/**
   声明有缓存的channel 不需要同时定义读写
 */

/**
 	写入三个数据到
 */
func CacheWrit(ch chan int){
	ch<-200
	fmt.Printf("ch addr：%v\n", ch) // 输出内存地址
	ch<-300
	fmt.Printf("ch addr：%v\n", ch) // 输出内存地址
   	ch<-400
	fmt.Printf("ch addr：%v\n", ch) // 输出内存地址

}
func main() {
   var ch chan int   // 声明一个有缓冲的channel
   ch=make(chan int,4)// 可以写入2个数据
   go CacheWrit(ch)
	// 防止主go程提前退出，导致其他协程未完成任务
	time.Sleep(time.Second * 3)

}

//总结 无缓冲通道与有缓冲通道
//无缓冲channel：
//
//通道的容量为0，即 cap(ch) = 0
//通道的个数为0，即 len(ch) = 0
//可以让读、写两端具备并发同步的能力
//有缓冲channel：
//
//在make创建的时候设置非0的容量值
//通道的个数为当前实际存储的数据个数
//缓冲区具备数据存储的能力，到达存储上限后才会阻塞，相当于具备了异步的能力
//有缓冲channel的阻塞产生条件：
//当缓冲通道被填满时，尝试再次发送数据会发生阻塞
//当缓冲通道为空时，尝试接收数据会发生阻塞
//问题：为什么 Go语言对通道要限制长度而不提供无限长度的通道?
//
//channel是在两个 goroutine 间通信的桥梁。使用 goroutine 的代码必然有一方提供数据，一方消费数据 。通道如果不限制长度，在生产速度大于消费速度时，内存将不断膨胀直到应用崩溃。
