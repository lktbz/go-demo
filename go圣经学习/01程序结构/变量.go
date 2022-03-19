package main

import (
	"fmt"
)

func main() {
	//单个变量声明
	var s string
	fmt.Println(s)
	//多个同类型的声明
	var i ,j ,k ,l int
	fmt.Println(i,j,k,l)
	//类型推导
	var b ,f ,ss,c= true,2.3,"four",false
	fmt.Println(b,f,ss,c)

	//简短变量声明
	 is:=100
	 bs:="string"
	 cs:=false
	 fmt.Println(is,bs,cs)
	//简短变量声明
	ii,jj:=0,1
	fmt.Println(ii,jj)

	//值的交换
	ii,jj=jj,ii
	fmt.Println(ii,jj)

}
