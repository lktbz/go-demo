package main

import "fmt"

type Persons struct {
	Username string
	Codes []int
}
func main() {
   p0:=&Persons{
   	"Davides",
   	[]int{31,32,35},
   }
   fmt.Println(p0)


//结构体切片
  p1:=[]Persons{
  	{"Nick",[]int{1,3,5}},
  	{"NiBi",[]int{2,4,6}},
  }
  fmt.Println(p1)


  //指针
  p2:=[]*Persons{
	  &Persons{"Nick",[]int{1,3,5}},
	  &Persons{"Ns",[]int{2,4,6}},
  }
  fmt.Println(p2)

  fmt.Println(*p2[0])

	for index, object := range p1 {
		fmt.Printf("object index:%v |username:%v|codes:%v |\n",
			index,object.Username,object.Codes)
	}

	for index, object := range p2 {
		fmt.Printf("object index:%v |username:%v|codes:%v |\n",
			index,object.Username,object.Codes)
	}

   //更新数据
   p1[0].Codes[0]=11
   fmt.Println("p1code:",p1[0].Codes)
   p2[0].Codes[0]=22
   fmt.Println("p2code:",p2[0].Codes)

}



