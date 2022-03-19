package main

import "fmt"

/**
   方法重写
 */

type Hum struct {
	name  string
	age   int
	phone string
}
type Stud struct {
	Hum  //匿名字段
	school string
}
type Empl struct {
	Hum   //匿名字段
	company string
}
//Hum定义method
func (h *Hum) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
func(e *Empl)SayHi()  {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}
func main() {
	mark := Stud{Hum{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Empl{Hum{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}
