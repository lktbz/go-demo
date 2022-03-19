package main

import (
	"fmt"
	"time"
)
//结构体
type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}

func main() {
	var dibert Employee
	//访问成员变量
	dibert.Salary=5000
	fmt.Println(dibert)
	//使用指针方式访问
	position:=&dibert.Position
	*position="senior"+*position

}
