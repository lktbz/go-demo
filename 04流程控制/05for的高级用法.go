package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName string
}

func main() {
	p1 := []Person{
		{"Nick", "Hill"},
		{"Facec", "Alex"},
	}
	fmt.Println(p1)

	for i := 0; i < len(p1); i++ {
		fmt.Println(p1[i])
		fmt.Println(p1[i].LastName, "--", p1[i].FirstName)
	}
	for index, obj := range p1 {
		fmt.Println(index,"==",obj)
	}



	m1:=map[string]int{
		"mark":10,
		"sandy":20,
		"Nick":5,
	}
	i:=0
	for key, value := range m1 {
		fmt.Println(i,key,value)
		i++
	}


}

