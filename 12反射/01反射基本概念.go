package main

import (
	"fmt"
	"reflect"
)
//首先必须懂反射
/**
reflect包的两个函数
reflect 提供了2个重要函数：

ValueOf()：获取变量的值，即pair中的 value
TypeOf()：获取变量的类型，即pair中的 concrete type
 */
type Person struct {
	Name string
	Age int
}

func main() {
	p:=Person{"lisi",23}
	//使用反射概念
	fmt.Println(reflect.ValueOf(p))//{lisi 23}
	fmt.Println(reflect.ValueOf(p).Type())//main.Person
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.TypeOf(p).Name())// Person:变量类型对象的类型名
	fmt.Println(reflect.TypeOf(p).Kind())	// struct:变量类型对象的种类名

	fmt.Println(reflect.TypeOf(p).Name() == "Person")	// true
	fmt.Println(reflect.TypeOf(p).Kind() == reflect.Struct)	//true

	/**
	类型与种类的区别：
	Type是原生数据类型： int、string、bool、float32 ，以及 type 定义的类型，对应的反射获取方法是 reflect.Type 中 的 Name()
	Kind是对象归属的品种：Int、Bool、Float32、Chan、String、Struct、Ptr（指针）、Map、Interface、Fune、Array、Slice、Unsafe Pointer等
	 */
	//静态类型与动态类型
	//静态类型：变量声明时候赋予的类型

	//type MyInt int			// int 是静态类型
	//var i *int				// *int 是静态类型
	//动态类型：运行时给这个变量赋值时，这个值的类型即为动态类型（为nil时没有动态类型）。

	//var A interface{}		// 空接口 是静态类型，必须是接口类型才能实现类型动态变化
	//A = 10					// 此时静态类型为 interface{} 动态为int
	//A = "hello"				// 此时静态类型为 interface{} 动态为string


}
