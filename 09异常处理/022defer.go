package main

import "fmt"

/**
   defer 延迟执行
 */
func main() {
   defer fmt.Println("lala")
   defer fmt.Println("kiki")
   fmt.Println("main...")


    num:=0
    defer  fmt.Println("defer 中的num=",num)
    num=3
    fmt.Println("main 中的num=",num)
}
