package main

import "fmt"

/**
     什么是结构体

  Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
  结构体是由一系列具有相同类型或不同类型的数据构成的数据集合
 */
//语法如下
/**
type struct_variable_type struct {
   member definition;
   member definition;
   ...
   member definition;
}
 */
//声明结构体
type person struct {
   age int
   name string
}

func main() {
	//结构体的初始化方式一：
	P:=person{age: 20,name: "lktbz"}
	fmt.Println(P)
	//方式二：

	P2:=person{22,"lktbz"}
	fmt.Println(P2)

	//方式三
	P3:= new(person)
	P3.name="jiji"
	P3.age=18
	//方式四
	p4:=&person{}

    fmt.Println(P3)
    fmt.Println(&P3)
    fmt.Println(P3.age,P3.name)
	fmt.Println(p4)
}
