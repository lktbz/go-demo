package main

import (
	"fmt"
	"sync"
)
/**
  读取跟期望的还是有差别
 */
var ww=sync.WaitGroup{}
var counters=0
var m =sync.RWMutex{}
func main() {
	//runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		ww.Add(2)
		go syaHello()
		go incre()
	}
	ww.Wait()

}

func incre() {
  m.Lock()
	counters++
  m.Unlock()
  ww.Done()
}

func syaHello() {
	m.RLock()
	fmt.Println("hello:=",counters)
	m.RUnlock()
	ww.Done()
}

