package main

import "fmt"

/**
    结构体构造修改
 */
type Empls2 struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}
//创建一个新函数包装一下构造
//func New(firstName string,lastName string,totalLeaves int,leavesTaken int) Empls2  {
//	e:=Empls2{firstName, lastName, totalLeaves, leavesTaken,}
//	return e
//}
//使用此方法调用。
func NewEmpls2(firstName string,lastName string,totalLeaves int,leavesTaken int)*Empls2{
	return &Empls2{firstName: firstName,lastName: lastName,totalLeaves: totalLeaves,leavesTaken: leavesTaken}
}
func (e Empls2) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}
func main() {
	//改了？
   //es:=Empls2.News("Sam", "Adolf", 30, 20)
	//es.LeavesRemaining() //有很大的问题?
	es:=NewEmpls2("Sam", "Adolf", 30, 20)
	es.LeavesRemaining()
}