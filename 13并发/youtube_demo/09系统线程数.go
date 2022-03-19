package main

import (
	"fmt"
	"runtime"
)


func main() {
   fmt.Println("线程数为：",runtime.GOMAXPROCS(-1))
}


