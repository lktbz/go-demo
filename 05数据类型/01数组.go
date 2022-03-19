package main

import (
	"fmt"
)

func main() {
	//声明语法
	//指明数组名字，元素个数，元素长度
	var names [10]int
	fmt.Println("定义的空names:", names)
	names[0] = 1
	names[1] = 2
	names[2] = 3
	fmt.Println("定义的names的三个值为:", names)
	//原始定义方式
	var arr1 = [10]string{}
	fmt.Println("arr1", arr1)
	var arr2 = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println("arr2", arr2)
	var arr3 = [5]int{'a', 'b', 'c', 'd', 'e'}
	fmt.Println("arr3", arr3)
	//简化操作
	arr4 := [5]int{4: 100}
	fmt.Println(arr4)
	arr5 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr5)
	f := [...]int{0: 1, 4: 1, 9: 1} // [1 0 0 0 1 0 0 0 0 1]
	fmt.Println(f)
	//数组的固定获取方式
	fmt.Println("arr3", arr2[2])
	fmt.Println("arr3", arr2[0:2])
	fmt.Println("arr3", arr2[1:3])
	fmt.Println("arr3", arr2[0:3])
	//遍历
	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("a", len(arr))
	//使用range 遍历数组
	for i, v := range arr {
		fmt.Println("元素的位置为", i)
		fmt.Println("元素的值", v)
	}
	/**
	  1.4 数组使用注意事项
	  数组创建完长度就固定，不可以再追加元素；
	  长度是数组类型的一部分，因此[3]int与[4]int是不同的类型；
	  数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该函数的副本，而不是他的指针。
	*/
	//值传递
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}
