package main

import "fmt"

/**
Go接口的特点
在上述示例中，Go无须像Java那样显式声明实现了哪个接口，即为非侵入式，接口编写者无需知道接口被哪些类型实现，接口实现者只需要知道实现的是什么样子的接口，但无需指明实现了哪个接口。编译器知道最终编译时使用哪个类型实现哪个接口，或者接口应该由谁来实现。

类型和接口之间有一对多和多对一的关系，即：

一个类型可以实现多个接口，接口间是彼此独立的，互相不知道对方的实现
多个类型也可以实现相同的接口
*/
//游戏服务器提供的服务
type service interface {
	//启动服务
	start()
	//日志服务
	log(string)
}

//日志器
type Logger struct {
}

//日志输出方法
func (log Logger) Log(s string) {
	fmt.Println("日志:=>", s)
}

//  游戏服务包含日志
type GameService struct {
	Logger
}

// 实现游戏服务的Start方法
func (game GameService) start() {
	fmt.Println("游戏服务启动。。。")
}

func main() {
	s := new(GameService)
	s.start()
	s.Log("hello game ")
}
