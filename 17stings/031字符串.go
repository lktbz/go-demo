package main

import (
	"fmt"
	"strings"
)

/**
strings 包函数学习
 */
func main() {
	s1:="helloworld"
   fmt.Println(strings.Contains(s1,"abc"))

   fmt.Println(strings.ContainsAny(s1,"abcd"))
	fmt.Println(strings.Count(s1,"lloo"))
	s2 := "20190525课堂笔记.txt"
   if strings.HasPrefix(s2,"201905"){
    fmt.Println("19年5月的文件")
   }

   if strings.HasSuffix(s2,".txt"){
	   fmt.Println("文本文档。。")
   }

    fmt.Println(strings.Index(s1,"lloo"))
    fmt.Println(strings.IndexAny(s1,"abcdh"))
	fmt.Println(strings.LastIndex(s1, "l"))    //查找substr在s中最后一次出现的位置

	//字符串的拼接
	ss1 := []string{"abc","world","hello","ruby"}
	s3:=strings.Join(ss1,".")
	fmt.Println(s3)

	s4 := "123,4563,aaa,49595,45"
	s5:=strings.Split(s4,",")
	fmt.Println(s5)
	//重复拼接字符串
	s6 :=strings.Repeat("hello",5)
	fmt.Println(s6)
	//n  为替换的 个数
	s7:=strings.Replace(s1,"l","*",1)
	s8:=strings.Replace(s1,"l","*",2)
	s9:=strings.Replace(s1,"l","*",-1)
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)

	s71:="heLLo WOrlD**123.."
	fmt.Println(strings.ToLower(s71))
	fmt.Println(strings.ToUpper(s71))
	fmt.Println(s1)
	/*
		截取子串：
		substring(start,end)-->substr
		str[start:end]-->substr
			包含start，不包含end下标
	*/
	s81 := s1[:5]
	fmt.Println(s81)
	fmt.Println(s1[5:])


}
