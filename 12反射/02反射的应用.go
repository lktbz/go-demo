package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age int
}
/**
    反射先了解吧，具体api 以后在分析
 */
func main() {
 //demo01()
	//demo02()
}
func demo02()  {
	//反射进行类型推断
	u := &User{
		Name: "Ruyue",
		Age:  100,
	}
	fmt.Println(reflect.TypeOf(u))
	fmt.Println(reflect.TypeOf(*u))
	fmt.Println(reflect.TypeOf(*u).Name())
	fmt.Println(reflect.TypeOf(*u).Kind())

}


func demo01(){
	/**
	反射操作简单数据类型
	 */
	// 设置值：指针传递
	var num int64 = 100
	ptrValue:=reflect.ValueOf(&num)
	newValue:=ptrValue.Elem()
	fmt.Println(newValue.Type())
	fmt.Println(newValue.CanSet())
	newValue.SetInt(230)
	fmt.Println(newValue)

	//值传递
	nums:=reflect.ValueOf(num)
	fmt.Println(nums.Int())
	fmt.Println(nums.Interface())
}