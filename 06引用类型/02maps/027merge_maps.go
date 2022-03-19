package main

import "fmt"

func main() {
first:=map[string]int{
	"a":1,
	"b":2,
	"c":3,
}
second:=map[string]int{
	"a":13,
	"e":5,
	"c":3,
	"d":4,
}

fmt.Println("============================")
	for i, v := range first {
		fmt.Printf("%v>%v",i,v)
	}
	for k, v := range second {
		first[k]=v
	}
	fmt.Println(first)



}
