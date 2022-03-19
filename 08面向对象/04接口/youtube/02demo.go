package main

import "fmt"

/**
   定义接口
 */
type Incrementer interface {
	Incrementer()int
}
//int 类型
type InCounter int

/**
  定义方法
 */
func (ic *InCounter)Incrementer()int  {
  *ic++
  return int(*ic)
}
func main() {
	 myInt:=InCounter(0)
	 var inc Incrementer=&myInt
	for i := 0; i < 10; i++ {
			fmt.Println(inc.Incrementer())
	}
}
