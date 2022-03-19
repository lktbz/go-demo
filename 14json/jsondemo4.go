package main

import (
	"encoding/json"
	"fmt"
)

/**

  map 转换成json
 */
func main() {
	t1:=make(map[string]interface{})
	t1["company"]="itDemo"
	t1["subjects "] = []string{"Go", "C++", "Python", "Test"}
	t1["isok"] = true
	t1["price"] = 666.666

	marshal, err := json.Marshal(t1)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(marshal))
}
