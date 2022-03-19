package main

import "fmt"


/**

    展示函数调用
 */

func main() {
	i, err := Add(20, 0)
	if err != nil {fmt.Println("失败了")
		return
	}
	fmt.Println("计算的结果为：",i)
}
