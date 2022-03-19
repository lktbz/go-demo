package main

import (
	"fmt"
)

func main() {
 	//sayGreeting("hello","lktbz")
 	greeting:="hello"
 	name:="stacey"
 	SayGreeting(&greeting,&name)
 	fmt.Println("3,",name)
}
//没有参数的函数
func sayGreeting(greeting,name string){ //简化参数类型
	fmt.Println(greeting,"--",name)
}

//指针方式，而不是副本方式
func SayGreeting(greeting,name *string){
	fmt.Println("1,",*greeting,*name)
	*name="TED"
	fmt.Println("2,",*name)
}