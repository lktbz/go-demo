package main

import (
	"fmt"
	"strconv"
)

/**

   字符串的类型转换
 */
func main() {
	s1:="true"
	e1, err := strconv.ParseBool(s1)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(e1)
	fmt.Printf("%T,%t\n",e1,e1)

	ss1 := strconv.FormatBool(e1)
	fmt.Printf("%T,%s\n",ss1,ss1)


	//2.整数
	s2 := "100"
	i2,err := strconv.ParseInt(s2,2,64)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n",i2,i2)

	ss2 := strconv.FormatInt(i2,10)
	fmt.Printf("%T,%s\n",ss2,ss2)
}
