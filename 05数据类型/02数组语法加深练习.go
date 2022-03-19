package main

import "fmt"

/**
  好记性比如烂笔头系列 数组语法加深学习
 */
func main() {
	//数组定义
	var ss=[5]int{1,2,3,4,5}
	fmt.Println(ss)
	var emp_array[5]int
	fmt.Println(emp_array)
	//数组开辟的空间
	fmt.Printf("%p\n",&emp_array)
	emp_array[0]=1
	emp_array[1]=2
	emp_array[2]=3
	emp_array[3]=4


	//打印第一个元素
	fmt.Println(emp_array[0])
	//打印第二个元素
	fmt.Println(emp_array[1])

    fmt.Println("数组的长度为:",len(emp_array))
    fmt.Println("数组的容量为:",cap(emp_array))


	//数组的其他创建方式
	var b=[5]int{1,2,3,4,5}
	fmt.Println(b)
	//在数组位置 2和4分别插入2和4
	var c=[5]int{1:2,3:4}
	fmt.Println(c) //[0 2 0 4 0]

	var d=[5]string{"rose","王二狗","ruby"}
	fmt.Println(d)

	//数组的创建之简化操作
    e:=[5]int{1,2,34,4,50}
    fmt.Println(e)
    f:=[...]int{1,2,3,4,5}
    fmt.Println(len(f))
    g:=[...]int{1:3,4:78}
    fmt.Println(g)


	//数组的遍历操作
	for i := 0; i < len(e); i++ {
      fmt.Println("i数组的值为：",e[i],"数组的下标为：",i)
	}

	for i, v := range e {
		fmt.Println("e数组的值为：",v,"数组的下标为：",i)
	}

	//只保留v 不要i
	for _,v := range e {
		fmt.Println("数组的值为：",v)
	}

}
