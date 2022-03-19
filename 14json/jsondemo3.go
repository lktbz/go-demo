package main

import (
	"encoding/json"
	"fmt"
)
/**

 控制字段输出样式以及是否在字段中显示
 */
type ITS struct {
	Company  string `json:"-"`   //添加- 此字段将会被忽略打印
	Subjects []string `json:"subjects"`  // Subjects 的值会进行二次JSON编码
	IsOk     bool	  `json:",string"` //转换为字符串，再输出
	Price    float64   `json:"price, omitempty"` // 如果 Price 为空，则不输出到JSON串中
}

func main() {
	t1 := ITS{Company: "itDance", Subjects: []string{"Go", "C++", "Python", "Test"}, IsOk: true}
	//格式化
	indent, _ := json.Marshal(t1)
	fmt.Println(string(indent))


}
