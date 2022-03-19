package main

import (
	"encoding/json"
	"fmt"


)



/**

  json  解析到interface
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

}
