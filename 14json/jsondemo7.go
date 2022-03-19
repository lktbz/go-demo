package main

import (
	"encoding/json"
	"fmt"


)
/**

  json使用断言判断类型
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
	var t interface{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(t)
	//使用断言判断类型
	m := t.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case bool:
			fmt.Println(k, "is bool", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

}
