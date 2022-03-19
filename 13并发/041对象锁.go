package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  对象锁
 */
type Account struct {
	money int
	lock  *sync.Mutex
}
func (a *Account)Query(){
  fmt.Println("当前金额为：",a.money)

}
func (a *Account)Add(num int){
	a.lock.Lock()
	a.money+=num
	a.lock.Unlock()
}
func main() {

	 a:=Account{0,&sync.Mutex{}}
	for i := 0; i <100 ; i++ {
		go func(num int) {
			a.Add(num)
		}(10)
	}
	time.Sleep(time.Second*2)
	a.Query()
}
