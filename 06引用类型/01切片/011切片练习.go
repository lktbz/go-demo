package main

import (
	"fmt"
)

func main() {
	/**
	  slice 就是可变数组，底层还是数组，解决数组存在的定长问题
	*/

	arr := [5]int{1, 2, 34, 56, 78}
	fmt.Println("这是数组", arr)
	arrtest := [4]int{1, 2, 13, 14}
	fmt.Println("切片的定义方式", arrtest)
	var s1 []int
	fmt.Println(s1)
	var s122 []int

	fmt.Println("这是空切片的定义", s122)

	s2 := []int{25, 34, 322, 1, 33}
	fmt.Println(s2)
	fmt.Printf("%T,%T\n", arr, s2) //[4]int,[]int

	//创建切片
	s3 := make([]int, 10, 10)
	fmt.Println(s3)
	s333 := make([]int, 5, 5)
	fmt.Printf("容量：%d,长度：%d\n", cap(s3), len(s3))
	fmt.Println("s333切片的方式为：", s333)
	s3[0] = 1
	s3[1] = 2
	s3[2] = 3
	fmt.Println(s3)
	//遍历添加数据
	s444 := make([]int, 5, 10)
	for i := 0; i < len(s444); i++ {
		s444[i] = i
	}
	fmt.Println("s444 的值为", s444)
	s444 = append(s444, 20, 50, 30)
	//s444=append(s444,"lktbz") 必须是相同类型
	s444 = append(s444, 99)
	fmt.Println("使用append 添加s444的值为：", s444)
	//报错。添加格式错误
	//s3[3] = 4 //panic: runtime error: index out of range [3] with length 3

	s3 = append(s3, 4, 5, 6)
	fmt.Println(s3, "长度为：", len(s3), "容量为：", cap(s3))
	s4 := make([]int, 0, 5)
	fmt.Println(s4)
	s4 = append(s4, 1, 2)
	fmt.Println(s4)
	s4 = append(s4, 3, 4, 5, 6, 7)
	fmt.Println(s4)
	s4 = append(s4, s3...)
	fmt.Println(s4)

	//遍历
	for i := 0; i < len(s4); i++ {
		fmt.Printf("%d", s4[i])
	}

	for i, v := range s4 {
		fmt.Println(i, v)

	}
	for i := 0; i < len(s444); i++ {
		fmt.Println("s444遍历方式一：", s444[i])
	}
	//i 为下标，v 为值
	for i, v := range s444 {
		fmt.Println("s444遍历方式二：", i, "---", v)
	}
	//从已有数组上创建切片
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("----------1.已有数组直接创建切片--------------------")
	s11 := a[:5]  //元素5个长度
	s22 := a[3:5] //元素长度为5，截取从第4个截取
	s33 := a[4:]  //从元素的第五个开始截取 ，一直到结束
	s44 := a[:]   //完整的将数组变成切片
	s11a := a[:6]
	fmt.Println("a:", a)
	fmt.Println("s11:", s11)
	fmt.Println("s22:", s22)
	fmt.Println("s33:", s33)
	fmt.Println("s44:", s44)
	fmt.Println("s11a", s11a)
	fmt.Println("----------3.更改数组的内容--------------------")
	a[4] = 100
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println("----------4.更改切片的内容--------------------")
	s2[2] = 200
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println("----------4.更改切片的内容--------------------")
	s1 = append(s1, 1, 1, 1, 1)
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println("----------5.添加元素切片扩容--------------------")
	fmt.Println(len(s1), cap(s1))

	s1 = append(s1, 2, 2, 2, 2, 2)
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(len(s1), cap(s1))
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", &a)
	fmt.Println(s1)
	//元素的删除
	s1 = append(s1[:1], s1[3:]...)
	fmt.Println(s1)


	var as=[10]int{1,2,3,4,5}
	fmt.Println("as---",as[:])
	fmt.Println(len(as))
}
