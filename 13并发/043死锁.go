package main

import (
	"fmt"
	"time"
)

/**
    死锁
 */
func main()  {
	//var rwm sync.RWMutex
	ch:=make(chan int)
	go func() {
		//rwm.RLock() //读锁
		x:=<-ch
		fmt.Println("读到的值为：",x)
		//rwm.RUnlock()

	}()
	go func() {
		  //rwm.Lock()
		  ch<-10
		  fmt.Println("写入：",10)
		  //rwm.Unlock()
	}()
	time.Sleep(time.Second*5)
}
