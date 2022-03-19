package main

import (
	"fmt"
	"runtime"
	"sync"
)
/**
 读取与期望一致
 */
var www=sync.WaitGroup{}
var counterss=0
var mm =sync.RWMutex{}
func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		www.Add(2)
		mm.RLock()
		go syaHellos()
		mm.Lock()
		go incres()
	}
	www.Wait()

}
func incres() {
	counterss++
	mm.Unlock()
	www.Done()
}
func syaHellos() {
	fmt.Println("hello:=",counterss)
	mm.RUnlock()
	www.Done()
}

