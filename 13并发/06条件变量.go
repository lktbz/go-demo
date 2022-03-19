package main

import (
	"fmt"
	"sync"
	"time"
)

/**
sync.Cond类型即是Go中的条件变量，该类型内部包含一个锁接口。
  条件变量通常与锁配合使用
*sync.Cond类型有三个方法：

Wait: 该方法会阻塞等待条件变量满足条件。也会对锁进行解锁，一旦收到通知则唤醒，并立即锁定该锁
Signal: 发送通知(单发)，给一个正在等待在该条件变量上的协程发送通知
Broadcast: 发送通知(广播），给正在等待该条件变量的所有协程发送通知，容易产生 惊群
 */
func main() {
   //创建并传入锁
	cond:=sync.NewCond(&sync.Mutex{})
	condition:=false
	// 开启一个新的协程，修改变量 condition
	go func() {
		time.Sleep(time.Second*1)
		cond.L.Lock()
		fmt.Println("准备更改消息")
		condition=true // 状态变更，发送通知
		cond.Signal()  // 发信号
		cond.L.Unlock()
	}()
	//main协程 是被通知的对象，等待通知
	cond.L.Lock()
	for !condition{
		  cond.Wait()
		  fmt.Println("获取到了消息")
	}
	cond.L.Unlock()
	fmt.Println("运行结束")

}
