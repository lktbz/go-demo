package main

import "fmt"

/**
   方法继承
 */

type Hu struct {
	name  string
	age   int
	phone string
}
type St struct {
	Hu  //匿名字段
	school string
}
type Em struct {
	Hu   //匿名字段
	company string
}

func (h *Hu) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
func main() {
	mark := St{Hu{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Em{Hu{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}
