package main

import "fmt"

/**
    函数之闭包
 */
func main() {
	/**
	1.1 闭包概念
	闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使己经离开了自由变量的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量。

	简单的说 : 函数+引用环境=闭包

	贴士：闭包( Closure)在某些编程语言中也被称为 Lambda表达式（如Java）

	在闭包中可以修改引用的变量
	 */

    str:="hello"
    foo:= func() {  // 声明一个匿名函数
		str="world"
	}
   foo()  //  调用匿名函数，修改str值
   fmt.Println(str) // world


	f := fn11(1)			//输出地址
	g := fn11(2)			//输出地址
	fmt.Println(f(1))		//输出1
	fmt.Println(f(1))		//输出1

	fmt.Println(g(2))		//输出2
	fmt.Println(g(2))		//输出2
	fmt.Println("-----------------------------")
	accAdd := Accumulate(1)
	fmt.Println(accAdd())				// 2
	fmt.Println(accAdd())
	fmt.Println(accAdd())
	fmt.Println(accAdd())
	fmt.Println(accAdd())
}
func fn11(a int)func(i int) int{
  return func(i int) int {
	    print(&a,"kaka",a)
		return a
  }
}
/**
   累加器
 */
func Accumulate(value int) func() int {
	return func() int {				// 返回一个闭包
		value++
		return value
	}
}