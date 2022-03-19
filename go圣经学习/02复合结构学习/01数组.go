package main

import "fmt"

func main() {
	 var a[3]int
	 fmt.Println(a[0])
	 fmt.Println(len(a)-1)
	// Print the indices and elements
	for i, v := range a {
		fmt.Printf("%d %d\n",i,v)
	}
	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}


	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2])
	fmt.Println(q)

	//...动态确认数组大小
	qs := [...]int{1, 2, 3}
	fmt.Printf("%T\n", qs)

}
