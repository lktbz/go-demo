package main

import (
	"fmt"
)

type Staff struct {
	Name,
	Language string
}
func main() {
s1,s2:=Staff{"dav","go"},Staff{"Alex","c++"}
  m1:=map[Staff]int{
  	s1:300,
  	s2:400,
  }
  fmt.Println(m1)

  varl,_:=m1[s1]
	fmt.Println(varl)
    _,check:=m1[s2]
	fmt.Println(check)
   m1[s1]=500
   fmt.Println(m1)
}
