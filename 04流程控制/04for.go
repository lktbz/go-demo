package main

import (
	"fmt"
)

func main() {
 result,limit:=0,100
 for first,second,third:=0,0,1;
   first<=limit;
   first,second,third=first+1,first+second,first+third{
   	if first==limit{
		result=first+second+third
	}
 }
 fmt.Println(result)

	for f:=1;f<=3;f++ {
		fmt.Println("===================")
		for m:=1;m<=3;m++ {
			fmt.Printf("%v,%v",f,m)

		}
	}




}
