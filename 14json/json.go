package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age int
}
/**  
   别名
 */
type PersonType struct {
	Name string `json:"my_name"`
	Age int `json:"my_age"`
}
func main() {
	p:=Person{Name: "lisi",Age: 50,}
	marshal, err := json.Marshal(&p)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
	//别名字段
	p2:=new(PersonType)
	p2.Name="goaldjf"
	p2.Age=20
	bytes, err := json.Marshal(p2)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))

	//反序列化

	//strs :=  `{"Name":"lisi","Age":50}`
    var i interface{}
	json.Unmarshal(marshal,&i)
	fmt.Println(i)
}
