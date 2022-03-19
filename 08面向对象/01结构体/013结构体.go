package main

import "fmt"

/**
结构体之间的嵌套
*/

type Address struct {
	city, state string
}
type Person struct {
	name string
	age  int
	address Address  //嵌套结构体
}
func main() {
   var p Person
    p.name="seven"
    p.age=50
    p.address=Address{
    	city:"chichji",
		state: "gugu",
	}

	fmt.Println("Name",p.name)
	fmt.Println("age",p.age)
	fmt.Println("city",p.address.city)
	fmt.Println("state",p.address.state)
}
