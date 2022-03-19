package main

import (
	"fmt"
)

func main() {
  sum(1,2,3,4,5,6,7)
	i := sumreturn(1, 2, 3, 4, 5, 6)
	fmt.Println(i)
	is := sumreturnpoints(1, 2, 3, 4, 5, 6)
	fmt.Println(is)
	fmt.Println(&is)
	fmt.Println(*is)
}
//多参数 传递
func sum(values ...int){ //简化参数类型
   fmt.Println(values)
   result:=0
	for _, v := range values {
		result+=v
	}
	fmt.Println("和为：",result)
}

//返回值的函数
func sumreturn(values ...int)int{
	result:=0
	for _, v := range values {
		result+=v
	}
	return result
}
//返回值的函数
func sumreturnpoints(values ...int)*int{
	result:=0
	for _, v := range values {
		result+=v
	}
	return &result
}

