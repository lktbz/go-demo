package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  在开发场景中，经常遇到多处并发读取，一次并发写入的情况，Go为了方便这些操作，在互斥锁基础上，提供了读写锁操作。

读写锁即针对读写操作的互斥锁，简单来说，就是将数据设定为 写模式（只写）或者读模式（只读）。使用读写锁可以分别针对读操作和写操作进行锁定和解锁操作。

读写锁的访问控制规则与互斥锁有所不同：

写操作与读操作之间也是互斥的
读写锁控制下的多个写操作之间是互斥的，即一路写
多个读操作之间不存在互斥关系，即多路读
 */
func main(){
	var  rwm sync.RWMutex
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println("Try Lock reading i:", i)
			rwm.RLock()
			fmt.Println("Ready Lock reading i:", i)
			time.Sleep(time.Second * 2)
			fmt.Println("Try Unlock reading i: ", i)
			rwm.RUnlock()
			fmt.Println("Ready Unlock reading i:", i)
		}(i)
	}
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Try Lock writing ")
	rwm.Lock()
	fmt.Println("Ready Locked writing ")
}