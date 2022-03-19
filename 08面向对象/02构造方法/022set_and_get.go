package main

import "fmt"

type SName struct {
	Name string
}
func(u *SName) SetName(name string){
	u.Name=name
}
func (u SName)SetName2(name string){
	u.Name=name
}
func main()  {
	p1:=&SName{}
	fmt.Println(&p1)
	fmt.Println(*p1)
	p1.SetName("A413")
	fmt.Println(p1.Name)
	p1.SetName2("B413")
	fmt.Println(p1.Name)
}