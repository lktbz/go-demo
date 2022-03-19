package main

import (
	"fmt"
	"sync"
)

/**
  去掉时间等待方式
*/
//等待组
var wg = sync.WaitGroup{}

func main() {
	var msg = "hello"
	wg.Add(1) //加一
	go func(msg string) {
		fmt.Println(msg) //打印的结果为hell0
		wg.Done()  //减一
	}(msg)
	msg = "Goodbye"
	wg.Wait()  //等待
}
