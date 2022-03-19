package main

import "fmt"

/**
   结构体的匿名字段
 */

type Human struct {
	name string
	age int
	weight int
}
type Student struct {
	Human // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}

func main() {
		// 我们初始化一个学生
		mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
		//访问匿名函数的字段以及自己的字段
		fmt.Println("His name is ",mark.name)
		fmt.Println("His age is ",mark.age)
		fmt.Println("His weight is ",mark.weight)
		fmt.Println("His speciality is ",mark.speciality)

	    // 修改对应的备注信息
        mark.speciality="AI"
		fmt.Println("His  change speciality is ",mark.speciality)
        marks:=Student{
        	Human{"lktbz",20,110},"jiujiu...."}
        fmt.Println(marks)
        fmt.Println(marks.Human)
	/**
	可以使用"."的方式进行调用匿名字段中的属性值

	实际就是字段的继承

	其中可以将匿名字段理解为字段名和字段类型都是同一个

	基于上面的理解，所以可以mark.Human = Human{"Marcus", 55, 220}和mark.Human.age -= 1

	若存在匿名字段中的字段与非匿名字段名字相同，则最外层的优先访问，就近原则
	 */


}
