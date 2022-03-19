package main

import (
	"fmt"
	"sort"
)

func main() {
    //map 的遍历
	map1 := make(map[int]string)
	map1[1] = "红孩儿"
	map1[2] = "小钻风"
	map1[3] = "白骨精"
	map1[4] = "白素贞"
	map1[5] = "金角大王"
	map1[6] = "王二狗"
	//因为数无序的所以每次遍历的结果顺序不一样
	for k, v := range map1 {
		fmt.Println(k,v)
	}

	fmt.Println("----------------------")

	for i := 1; i <=len(map1) ; i++ {
       fmt.Println(i,"---",map1[i])
	}
	fmt.Println("----------------------")
	for i := 0; i <=len(map1) ; i++ {
		fmt.Println(i,"---",map1[i])
	}

	/*
		1.获取所有的key，-->切片/数组
		2.进行排序
		3.遍历key，--->map[key]
	*/
    keys:=make([]int ,0,len(map1))
     fmt.Println(keys)
	for i, _ := range map1 {
		keys = append(keys, i)
	}
     fmt.Println(keys)
     sort.Ints(keys)
	 fmt.Println(keys)


	values:=make([]string ,0,len(map1))
	fmt.Println(values)
	for _, v := range map1 {
		values= append(values, v)
	}
	fmt.Println(values)

	for _,key :=range keys{
		fmt.Println(key,map1[key])
	}

	s1 :=[]string{"Apple","Windows","Orange","abc","王二狗","acd","acc"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)


}
