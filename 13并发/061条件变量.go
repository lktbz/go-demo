package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//使用条件变量优化生产消费模型（支持多个生产者、多个消费者）

//缓冲区大小
const BUFLEN=5
//定义全局变量
var con *sync.Cond =sync.NewCond(&sync.Mutex{})
//定义生产者
func producers(ch chan <-int){
	for{
		con.L.Lock()// 给条件变量对应的互斥锁加锁
		for len(ch)==BUFLEN { // 缓冲区满，则等待消费者消费，这里不能是if
			con.Wait()
		}
		ch<-rand.Intn(1000)// 写入缓冲区一个随机数
		con.L.Unlock()// 生产结束，解锁互斥锁
		con.Signal() // 一旦生产后，就唤醒其他被阻塞的消费者
		time.Sleep(time.Second*2)
	}
}

func consumers(ch <-chan int ){
	for{
		con.L.Lock()

		for len(ch)==0 {
			con.Wait()
		}
		fmt.Println("接收：",<-ch)
		con.L.Unlock()
		con.Signal()
		time.Sleep(time.Second*1)
	}

}
func main() {
   rand.Seed(time.Now().UnixNano())
   ch:=make(chan int ,BUFLEN)
	for i := 0; i < 10; i++ {
		go producers(ch)
	}
	for i := 0; i <10 ; i++ {
		go consumers(ch)
	}
	for  {
		
	}
}
