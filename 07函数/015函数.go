package main

import (
	"errors"
	"fmt"
)
/**
   容错函数定义
 */
func Add(a int,b int)(ret int,err error){
	if a<0 ||b<0{
		errors.New("不能输入小于0的数")
		return
	}
	return a+b,nil
}
func main() {
	i, err := Add(10, 20)
	if err != nil {
		fmt.Println("函数运行失败")
		return
	}
	fmt.Println("函数运行的结果为：",i)
}
