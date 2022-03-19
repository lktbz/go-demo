package main

import (
	"fmt"
	"sort"
)

func main() {
  ages:=map[string]int{
  	    "alice":33,
  	    "charlies":34,
	}
 fmt.Println(ages)
  agess:=make(map[string]int)
  agess["alice"]=31

  agess["chars"]=32
  fmt.Println(agess)

  fmt.Println(agess["chars"])
  delete(agess,"alice")
  fmt.Println(agess)


  var names[] string
	for name := range ages {
		names=append(names,name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	var aages map[string]int
	fmt.Println(aages==nil)
	fmt.Println(len(aages)==0)
}
