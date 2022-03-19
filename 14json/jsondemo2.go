package main

import (
	"encoding/json"
	"fmt"
)
/**
  普通的json 没有对字段进行限制

 */
type IT struct {
	Company  string
	Subjects []string
	IsOk     bool
	Price    float64
}

func main() {
	t1 := IT{"itDance", []string{"GO", "C++", "Python", "Test"}, true, 66.66}
	marshal, _ := json.Marshal(t1)
	fmt.Println(string(marshal))
	//格式化
	indent, _ := json.MarshalIndent(t1, "", "  ")
	fmt.Println(string(indent))


	indent1, _ := json.MarshalIndent(t1, "lktbz", "  ")
	fmt.Println(string(indent1))

	indent2, _ := json.MarshalIndent(t1, "", "--")
	fmt.Println(string(indent2))
}
