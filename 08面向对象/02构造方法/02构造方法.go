package main

import "fmt"

/**
   构造
 */
//创建结构体pt
type Pt struct {
 	Name string
	Age int
}
//创建构造函数 参数为name
func newPersonByName(name string)*Pt  {
	  return &Pt{
			Name: name,
	  }
}
//创建构造函数，参数为age
func  newPersonByAge( age int)*Pt  {
	return &Pt{Age: age,
		}
}
//创建构造函数，参数为 name,age
func newPersonByNameAndAge(name string,age int)*Pt  {
    return &Pt{
    	Name: name,
    	Age: age,
	}
}
//构造函数方式
func newPersonDemo(name string, age int)*Pt{
	return &Pt{
		Name: name,
		Age: age,
	}
}

func main() {
	p1:=newPersonByName("lktbz")
	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
	p2:=newPersonByAge(20)
	fmt.Println(p2.Age)
	fmt.Println(p2.Name)
	p3:=newPersonByNameAndAge("juli",25)
	fmt.Println(p3.Name,"-",p3.Age)

}
