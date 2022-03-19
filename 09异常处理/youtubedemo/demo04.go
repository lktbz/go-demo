package main

import "fmt"
/**
    这会是什么运行结果
 */
func main() {
	fmt.Println("start")//执行
	defer fmt.Println("this was defered")  //执行
	panic("something bad happend")  //抛出错误，终止程序执行
	fmt.Println("end") //没有执行
}
