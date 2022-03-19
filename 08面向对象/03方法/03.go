package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

//带参数以及返回值的方法
func Distance(p,q Point) float64  {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	//直接方法调用
	fmt.Println(Distance(p, q)) //调用的第一个方法
	//类中方法调用
	fmt.Println(p.Distance(q)) //调用的第二个方法
}
