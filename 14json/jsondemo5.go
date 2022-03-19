package main

import (
	"encoding/json"
	"fmt"


)

type ITs struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}


/**

  json 转换成对象
 */
func main() {
	//首先转换成字节
	b := []byte(`{
    "company": "itcast",
    "subjects": [
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "isok": true,
    "price": 666.666
	}`)

	var t ITs
	err := json.Unmarshal(b, &t)
	if err!=nil{
		fmt.Println("解析失败。。。")
	}
	fmt.Println("解析的结果为：",t)
}
