package main

import "fmt"

/**
    使用匿名结构体，来简化访问
//下面这段话可以简单理解为， 引用 匿名结构体的结构体把匿名结构体中的字段全部copy 一份到自己的结构体中
在结构体中属于匿名结构体的字段称为提升字段，因为它们可以被访问，
就好像它们属于拥有匿名结构字段的结构一样。理解这个定义是相当复杂的。
 */

type Address1 struct {
	city, state string
}
type Person1 struct {
	name string
	age  int
	Address1
}
func main() {
	//p2 类型的person1 ,可以使用全部的类型
	var p2 Person1
	p2.age=20
	p2.name="ll"
	p2.state="CN"
	p2.city="bj"
	fmt.Println("Name:", p2.name)
	fmt.Println("Age:", p2.age)
	fmt.Println("City:", p2.city) //city is promoted field
	fmt.Println("State:", p2.state) //state is promoted field

}
