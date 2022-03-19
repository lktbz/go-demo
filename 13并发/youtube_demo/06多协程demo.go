package main

import (
	"fmt"
	"sync"
)

var wgg = sync.WaitGroup{}
var counter = 0
/**
   多协程工作，没有同步操作
 */
func main() {
	for i := 0; i < 10; i++ {
		wgg.Add(2)
		go sayHello2()
		go increment()
	}
	wgg.Wait()
}

func increment() {
	counter++
	wgg.Done()
}
//打印的结果一团槽
func sayHello2() {
	fmt.Printf("hello #%d\n", counter)
	wgg.Done()
}
