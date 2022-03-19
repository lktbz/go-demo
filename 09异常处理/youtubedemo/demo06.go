package main

import (
	"fmt"
	"log"
)
//虽然函数抛出异常了，但是我们还是需要执行相关方法
func main() {
 fmt.Println("start")
 panicker()
 fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err:=recover();err!=nil{
			log.Println("error:",err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}
