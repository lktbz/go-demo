package main

import "fmt"

func main() {
	const b  string = "abc" //直接定义常量
	const  c  = 123 //隐式定义常量

	const  LENGTH int =10
	const WIDTH int =5
	var area int
	const aa, bb, cc =1,false,"lktbz"
	area=LENGTH*WIDTH
	fmt.Println("面积为：%d",area)
	println()
    println(aa,bb,cc)

   //常量池
	const(
		   UK=10
			UN=11
			CN=12
	)
    fmt.Println(UK,UN,CN)


	//iota，特殊常量，可以认为是一个可以被编译器修改的常量
	//const (
	//	ia = iota
	//	ib = iota
	//	ic = iota
	//	e ="a"
	//)
	//
	//fmt.Println(ia,ib,ic,e)

	fmt.Println()

	//累加器
	const (
		iq=iota
		iw=iota
		ie=iota
		ir="j"
		it="k"
		iu=iota
		io
	)
	fmt.Println(iq,iw,ie,ir,it,iu,io)

}
