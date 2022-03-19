package main

import "fmt"

/**
  嵌套函数的初始化
 */
//Pubs看做父类，Subs 看做子类
type Pubs struct {
   Name string
   age int
}
type Subs struct {
	Pubs
	foodName string
}
//父类构造
func  newPubs(name string,age int)*Pubs  {
	return &Pubs{
		Name: name,
		age: age,
	}
}
//子类构造
func newSubs(foodName string)*Subs  {
	 p:=&Subs{
	 }
	 p.foodName=foodName
	 return p
}

func main() {
	s := newSubs("一哥")
	fmt.Println(s)

	//这种构造方式，感觉很是奇怪
	//直接使用这种方式进行替代
	sb:=new(Subs)
	sb.foodName="beef"
	sb.age=20
	sb.Name="black girls"
	fmt.Println(sb)

}
