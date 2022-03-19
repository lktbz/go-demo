package main

import (
	"fmt"
	"log"
)
/**
    这会是什么运行结果
 */
func main() {
	fmt.Println("start")//执行
	defer func() {
		//recover恢复函数，当发生，panic错误，err b等于nil
		if err:=recover();err!=nil{
			log.Println("Error:" ,err)
		}
	}()
	panic("something bad happend")  //抛出错误，终止程序执行
	fmt.Println("end") //没有执行
}
