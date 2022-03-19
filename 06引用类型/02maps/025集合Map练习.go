package main

import (
	"fmt"
)
/**
     map 中结构体的练习
 */
type PersonInfo struct {
	ID string
	Name string
	Address string
}

func main() {
	/**
	   声明方式分析：
	   personDB 变量名
	  string key的类型
		PersonInfo value 的类型
	 */
 var personDB map[string]PersonInfo
 personDB=make(map[string]PersonInfo)
 personDB["12345"]=PersonInfo{"1234","Tom","Room23"}
 personDB["1"]=PersonInfo{"1","Jack","Room101"}
  person,ok:=personDB["12345"]
   if ok {
   	fmt.Println(person)
   }else{
   	fmt.Println("出错了。。。。")
   }
   //初始包含100个map的能力
   var lenMap =make(map[string]string,100)
   fmt.Println(lenMap)
   fmt.Println("length:=",len(lenMap))

   //创建map 并进行初始化
   var initMap=map[string]string{"zs":"sd","to":"name","sa":"saName"}
   fmt.Println(initMap)
   fmt.Println("length:=",len(initMap))
}
