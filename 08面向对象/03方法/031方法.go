package main

import "fmt"

/**
 t Type
func (t Type) methodName(parameter list)(return list) {

}
func funcName(parameter list)(return list){

}
  方法的简单使用
 */

type Emp struct {
	name     string
	salary   int
	currency string
}

//定义一个打印的方法
//
func( e Emp)print(){
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}
func main() {
   mp:=Emp{
   	  name: "lktbz",
   	  salary: 2500,
   	  currency:"&",
   }
   //调用方法
   mp.print()
}
