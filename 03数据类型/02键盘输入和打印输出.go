package main

import "fmt"

func main() {
	var x int
	var y float64
	fmt.Println("请输入一个整数，一个浮点数")
	fmt.Scanln(&x, &y)
	fmt.Printf("x的数值%d,y的数值为%f", x, y)

	fmt.Scanf("%d,%f", &x, &y)
	fmt.Printf("x:%d,y:%f\n", x, y)

}

