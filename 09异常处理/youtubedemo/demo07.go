package main

import (
	"fmt"
	"log"
)
//优雅的处理了错误，但是还是出错就终止程序执行
func main() {
 fmt.Println("start")
 panicker1()
 fmt.Println("end")
}

func panicker1() {
	fmt.Println("about to panic")
	defer func() {
		if err:=recover();err!=nil{
			log.Println("error:",err)
			panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}
