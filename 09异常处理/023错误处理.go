package main

import "fmt"

/**
panic  运行到相应函数将阻止执行，defer 将会继续执行

阻止程序继续往下执行
 */

func  demo01()  {
 fmt.Println("执行。。。。")
 fmt.Println("执行22。。。。")
 fmt.Println("执行33。。。。")
 panic("执行错误。，。。")

 fmt.Println("执行334。。。。")
}
func  demo02()  {
 fmt.Println("fefer 调用。。。")
}
func main() {
 fmt.Println("即将执行。。。。。。")
   defer demo02()
   demo01()
   defer demo02()
}
