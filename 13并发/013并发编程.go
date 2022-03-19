package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
goroutine 常用api
 */
func main() {
  //cpuNums()
	//GoschedDemo()
	GoExit()
}
/**
Goexit() 终止当前协程
 */
func GoExit()  {
	for i := 1; i <= 5; i++ {
		defer fmt.Println("defer ", i)
		go func(i int){
			if i == 3 {
				runtime.Goexit()
			}
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
}
/***
   让出时间片
 */
func GoschedDemo()  {
	//可以理解为接力赛跑：A跑了一段遇到了Gosched接力给B。
	for i := 0; i < 10; i++ {
       go func(i int ) {
			if i==5{
  			runtime.Gosched()// 协程让出，但并不代表不执行，而是 5永远不会第一输出
			}
			fmt.Println(i)
	   }(i)
	}
	time.Sleep(time.Second)
}
func cpuNums()  {
	cpu:= runtime.NumCPU()
	fmt.Println("cpu 的核心数为",cpu)
}
