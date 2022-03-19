package main

import "fmt"

//接口 phone
type Phone interface {
	//方法
	call()
}
//定义结构体
type ApplePhone struct {
}

//实现接口方法
func (apple ApplePhone) call() {
	fmt.Println("I am Apple, I can call you!")
}
type HuaWei struct {
}
func (huawei HuaWei) call() {
	fmt.Println("I am HuaWei, I can call you!")
}
func main() {
	var phone Phone
	phone =new(ApplePhone)
	phone.call()

	phone = new(HuaWei)
	phone.call()
}
