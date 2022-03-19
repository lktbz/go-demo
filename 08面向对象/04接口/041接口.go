package main

import "fmt"

/**
		interface  值
 */
//先看代码
//接口定义
type Men interface {
	SayHi()
	Sing(lyrics string)
}
//结构体
type Humans struct {
	name string
	age int
	phone string
}
//嵌套结构体
type Students struct {
	Humans
	school string
    loan   float32
}
type Employees struct {
	Humans
	company string
    money   float32
}
//调用Men 接口的方法
func(hs Humans)SayHi(){
	fmt.Printf("Hi, I am %s you can call me on %s\n", hs.name, hs.phone)
}
//调用Men 接口的方法
func (hs Humans)Sing(lyrics  string)  {
	fmt.Println("La la la la...", lyrics)
}
//调用Men 接口的方法
func(e Employees)SayHi(){
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split i
}
/**
  使用指针方式
 */
//func (e *Employees)SayHi(){
//	 fmt.Println(e.name,e.phone)
//}


func main() {
	mike := Students{Humans{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Students{Humans{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employees{Humans{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	Tom := Employees{Humans{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}
	//定义Men类型的变量i
	var i Men
	//i能存储Student
	i=mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
 	i=Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")
	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
 	//定义数组里面继续调用，用变量指向对象？
	x:=make([]Men,3)
	x[0],x[1],x[2]=paul,sam,mike
	for _, men := range x {
    	  men.SayHi()
	}
	/**
	那么interface里面到底能存什么值呢？如果我们定义了一个interface的变量，那么这个变量里面可以存实现这个interface的任意类型的对象。例如上面例子中，我们定义了一个Men interface类型的变量m，那么m里面可以存Human、Student或者Employee值

	当然，使用指针的方式，也是可以的

	但是，接口对象不能调用实现对象的属性
	 */
}
