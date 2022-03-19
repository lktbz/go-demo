package main

import "fmt"

func main() {
  l1:=[3][3]int{
   [3]int{1,3,5},
   [3]int{2,4,6},
   [3]int{7,8,9},
  }
  fmt.Println(l1,len(l1))


  l2:=[3][2]float32{
  	[2]float32{1.2,3.4},
  	[2]float32{2.3,4.5},
  	[2]float32{3.1,5.6},
  }
  fmt.Println(l2)

  l1[0]=[3]int{11,33,55}
  fmt.Println(l1)


   l1[1][1]=55
   fmt.Println(l1)

	//多维切片
   s:=[][]int{
   	[]int{1,2,1,1},
   	[]int{1,2,1,1},
   	[]int{1,2},
   	[]int{3},
   }
   fmt.Println(s)
   s[1]=[]int{44}
   s[1]=append(s[1],55)
   fmt.Println(s)
}
