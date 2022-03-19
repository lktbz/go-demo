package main

import (
	"fmt"
	"math"
)
//x和y是形参名,3和4是调用时的传入的实数，函数返回了一个float64类型的值。
func hypot(x,y float64)float64{
	return  math.Sqrt(x*x+y*y)
}
//如果一组形参或返回值有相同的类型，我们不必为每个形参都写出参数类
//型。下面2个声明是等价的：
//func f(i, j, k int, s, t string) { /* ... */ }
func f(i int, j int, k int, s string, t string) { /* ... */ }

func main() {
fmt.Println(hypot(3,4))
}
