package main

import (
	"fmt"
)

type Users struct {
	userName string
	Conections map[string]int32
}
func main() {
  p0:=&Users{"Dc", map[string]int32{"fd":3572,"ray":2421}}
  fmt.Println(p0)


  p1:=[]Users{
  	{"nick", map[string]int32{"f":1,"sf":415}},
  	{"niubi", map[string]int32{"n":2,"no":4412}},
  }

  fmt.Println(p1)

	for index, object := range p1 {
		fmt.Printf("object index:%v |username:%v|conns:%v |\n",
			index,object.userName,object.Conections)
	}

//修改数据
p1[0].Conections["f"]=8431
fmt.Println(p1[0].Conections)
}
