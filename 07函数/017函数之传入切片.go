package main

import "fmt"

func SetUpConnections(m[]string,f[]int32)(map[string]int32,string){
	if len(m)!=len(f){
		return  nil,"长度不相等"
	}
	m1:=make(map[string]int32)
	for index, value := range m {
		m1[value]=f[index]
	}
	return m1,"已经创建"
}
func main()  {
	l1:=[]string{"davi","mid"}
	l2:=[]int32{23451,11121}
	connections, s := SetUpConnections(l1, l2)
	if connections!=nil{
		fmt.Println(connections,s)
	}else{
		fmt.Println(connections)
	}
}
