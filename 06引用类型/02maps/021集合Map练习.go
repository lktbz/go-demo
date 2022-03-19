package main

import "fmt"

func main() {
	//map的创建
   var mp1 map[int]string  //没有初始化，nil
   var mp2=make(map[string]string)
   var mp3=map[string]int{"Go": 98, "Python": 87, "Java": 79, "Html": 93}
   fmt.Println(mp1)
   fmt.Println(len(mp1))
   fmt.Println(mp2)
   fmt.Println(mp3)

   fmt.Println(mp1==nil)
   fmt.Println(mp2==nil)
   fmt.Println(mp3==nil)
   //初始化mp1
   if mp1==nil{
		mp1=make(map[int]string)
	    fmt.Println(mp1)
	}
	fmt.Println(mp1==nil)
	//添加数据
	mp1[1] = "hello" //panic: assignment to entry in nil map
	mp1[2] = "world"
	mp1[3] = "memeda"
	mp1[4] = "王二狗"
	mp1[5] = "ruby"
	mp1[6] = "三胖思密达"
	mp1[7] = ""
	fmt.Println("mp1",mp1)
	fmt.Println(mp1[4]) //根据key为4，获取对应的value值
	fmt.Println(mp1[40]) //操作不存在的值


     v1,ok:=mp1[40]
	if ok {
		fmt.Println("对应的数值是：",v1)
	}else{
		fmt.Println("操作的key不存在，获取到的是零值：",v1)
	}
	//修改数据
	fmt.Println(mp1)
	mp1[3]="如花"
	fmt.Println(mp1)
	//删除数据
	delete(mp1,3)
	fmt.Println(mp1)
	delete(mp1,30)
	fmt.Println(mp1)
	//长度
	fmt.Println(len(mp1))


}
