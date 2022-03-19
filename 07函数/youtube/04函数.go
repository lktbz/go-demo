package main

import "fmt"

func main() {
  //d:=divide(5.0,2.0) //一切那么美好
  //d:=divide(5.0,0.0) //传入0
  d,err:=divides(5.0,3.0) //传入0
	if err!=nil {
		fmt.Println("程序出错。。。")
		return
	}
  fmt.Println(d)
}
//函数的错误处理
func divide(a,b float64)float64  {
	//处理传入0 值的情况
	if b==0.0{
		panic("不能除以0") //这么处理，是真的不优雅
	}
	return a/b
}
/**
	返回错误
 */
func divides(a,b float64)(float64,error)  {
	if b==0.0{
		return 0.0, fmt.Errorf("不能输入0")
	}
	return a/b,nil
}