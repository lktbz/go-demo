package main

import "fmt"

func main() {
  var input int =0
  L1:
  	input++
  	if input<=10{
  		fmt.Print(">",input)
  		goto L1
	}else {
		fmt.Println("we left for loop!")
	}
}
