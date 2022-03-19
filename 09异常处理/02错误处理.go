package main

import (
	"errors"
	"fmt"
)
/**
    go的错误处理，，。。。go语言错误处理方式为往上抛让调用者去处理
 */
/**
  函数的多返回类型
 */
func add(a int ,b int)(n int, err error)  {
  if a<=0&&b<=0{
  	return 0,errors.New("非法的参数，参数必须大于0")
  }else {
  	sum:=a+b
	  return sum, nil
  }
}
func main() {
  //常见的调用可能异常的方式
  res,ok:=add(10,15)
  if ok==nil{
  	fmt.Println("结果为",res)
  }

}
