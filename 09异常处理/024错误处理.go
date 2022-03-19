package main

import "fmt"

/**
Go提供的recover机制：由运行时抛出，或者开发者主动触发的panic，可以使用defer和recover实现错误捕捉和处理，让代码在发生崩溃后允许继续执行。

在其他语言里，宕机往往以异常的形式存在，底层抛出异常，上层逻辑通过try/catch机制捕获异常，没有被捕获的严重异常会导致宕机，
捕获的异常可以被忽略，让代码继续执行。Go没有异常系统，使用panic触发宕机类似于其他语言的抛出异常，recover的宕机恢复机制就对应try/catch机制
panic和defer的组合：

有panic没有recover，程序宕机
有panic也有recover，程序不会宕机，执行完对应的defer后，从宕机点退出当前函数后继续执行

*/

func test(n1 int ,n2 int)  {
	defer func() {
		err:=recover()
		if err!=nil{
			fmt.Println("err=",err)
		}
	}()
	fmt.Println(n1*n2)
}
func main(){
  test(20,-1)
  fmt.Println("after....")

}
