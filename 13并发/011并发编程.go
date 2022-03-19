package main

import (
	"fmt"
	"time"
)

/**
    同时执行两件事
 */

func  running()  {
	var times int
	for {
		times++
		fmt.Println("tick:", times)
		time.Sleep(time.Second)
	}
}
func main() {
   go running()
    var input string
	fmt.Scanln(&input) //接受用户输入，返回并终止程序

}
