package main

import "fmt"

func main() {
/**
	指针是存储另一个变量的内存地址的变量。

  我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。

  一个指针变量可以指向任何一个值的内存地址它指向那个值的内存地址。
 */


  a:=10
  fmt.Println(&a)//取a 的地址
  fmt.Println(a)//取a的值
  //定义数组
  arr:=[5]int{1,2,3,4,5}
	for i := 0; i <len(arr) ; i++ {
		fmt.Printf("数组下标%d的内存地址为%d\n",i,&arr[i])
	}

  //指针的声明

   var ip *int
   ip=&a //指针变量ip 存储a 的地址引用
	fmt.Printf("a 变量的地址是: %x\n", &a  )

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量的存储地址: %x\n", ip )

	fmt.Printf("ip和a的两个是否相等%t\n",&a==ip)

	fmt.Printf("*ip变量的值为:%d\n",*ip)
	fmt.Printf("*ip变量的地址为:%d\n",ip)



	//获取指针的值
	cd:=255
    ae:=&cd
	fmt.Println("address of b is", ae)
	fmt.Println("value of b is", *ae)


	//操作指针改变变量的数值
	*ae++
	fmt.Println("new value of b is", cd)


    //使用指针传递函数的参数
	fmt.Println("value of a before function call is",cd)
    change(ae)
	fmt.Println("value of a after function call is", cd)


    //指针的指针
	var ptr *int
	var pptr **int
	/* 指针 ptr 地址 */
	ptr=&a
	/* 指向指针 ptr 地址 */
	pptr=&ptr
	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a )
	fmt.Printf("指针变量 *ptr = %d\n", *ptr )
	fmt.Printf("指针变量 ptr = %d\n", ptr )
	fmt.Printf("指针变量 &ptr = %d\n", &ptr )
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
	fmt.Printf("指向指针的指针变量 pptr = %d\n", pptr)
	fmt.Printf("指向指针的指针变量 &pptr = %d\n", &pptr)

  var ms *mystruct
	fmt.Println(ms)
	//ms=&mystruct{foo: 42}
   //fmt.Println("指针类型：",ms)
   //使用new 替换上面的
 ms=new(mystruct)
	(*ms).foo=42
 //ms.foo=22  这是方式是推荐的，简单理解
 //fmt.Println(ms)
 fmt.Println((*ms).foo)  //指针取值，没有地址符号
}

type mystruct struct {
	foo int
}
func change(val *int) {
	*val = 55
}