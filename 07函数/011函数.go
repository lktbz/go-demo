package main

import "fmt"

/**
   函数高级用法总结
 */
func main() {
	a:=5
	fmt.Println("===========niming函数实例================")
	//f1 指向匿名函数
	f1:= func(num int) {
		fmt.Println("匿名函数等待被调用：-",num)
	}
   f1(a)

	fmt.Println("===========自调用函数实例================")
   func(){
   	fmt.Println("匿名函数自调用",a)
   }()
	fmt.Println("===========匿名函数求最大值最小值================")
	x,y:= func(i,j int)(max,min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10,20) //自调用并传值
	fmt.Println(x + ' ' + y)

}
/**
函数去掉函数名、参数名和{}后的结果即是函数类型，可以使用%T打印该结果。

两个函数类型相同的前提是：拥有相同的形参列表和返回值列表，且列表元素的次序、类型都相同，形参名可以不同。
 */
func mathSum(a, b int) int {
	return a + b
}

func mathSub(a, b int) int {
	return a - b
}

//定义一个函数类型
type MyMath func(int, int) int

//定义的函数类型作为参数使用
func Test(f MyMath, a , b int) int{
	return f(a,b)
}
//通常可以把函数类型当做一种引用类型，实际函数类型变量和函数名都可以当做指针变量，只想函数代码开始的位置，没有初始化的函数默认值是nil。
