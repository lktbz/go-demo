package main

import "fmt"

func main() {
	a := "start"
	defer fmt.Println(a)
	a = "end"
}
