package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//普通加锁
   var mu sync.Mutex
   num:=0
	for i := 0; i <1000 ; i++ {
		 go func() {
			 mu.Lock()
			 num+=1
			 mu.Unlock()
		 }()
	}
	time.Sleep(time.Second*5)
	fmt.Println("num:",num)
}
