package main

import "fmt"

func main() {
  func(){
  	 msg:="hello go"
  	fmt.Println(msg)
  }()

  //匿名函数的自调用
	  for i := 0; i <5 ; i++ {
		  func(i int){ //接受的参数
		  	fmt.Println(i)
		  }(i) //实际传送的值
	  }

}
