package main

import "fmt"

func main() {
 name1,price:="ben",413
	switch name1 {
	case "mike":
		fmt.Println("我是 mike")
	case "ben":
		fmt.Println("我是ben")
	default:
		fmt.Println("我谁也不是。。。")
	}

	switch  {
	case price<413:
		fmt.Println("price 小于 413")
	case price>413:
		fmt.Println("price 大于413")
	default:
		fmt.Println("price 等于413")

	}

	name2:="ben"
	switch name2 {
	case "da","mike","ben":
		fmt.Println("name is da or mike or ben")
	default:
		fmt.Println("谁也不是")
	}

	switch newPrice:=500; {
	case newPrice<500:
		fmt.Println("price 小于500")
	case newPrice==500:
		fmt.Println("price 等于500")
	default:
		fmt.Println("price 大于500")

	}




	//条件不满足也会执行。。。。
	process:="Scheduler"
	switch process {
	case "Scheduler":
		fmt.Println("Scheduler")
		fallthrough
	case "wen":
		fmt.Println("wen")
		fallthrough
	case "fri":
		fmt.Println("fri")
		fallthrough
	default:
		fmt.Println("no....")
	}
}
