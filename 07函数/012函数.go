package main

import "fmt"

//init 函数
func init()  {
	fmt.Println("init 函数运行") //运行在前
}
func main() {

	fmt.Println("main 函数运行")


	/**
	new函数可以用来创建变量。表达式new(T)将创建一个T类型的匿名变量，
	初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T：
	 */

	p := new(int)		// p 为 *int类型，只想匿名的int变量
	fmt.Println(*p)		// "0"
	*p = 2				// 设置 int匿名变量值为2
	fmt.Println(*p)
}
