package main

import "fmt"
/**
 	函数当类型进行调用
 */
func main() {
	//声明函数类型
   var divide func(float64,float64)(float64,error)
   //实际的函数类型执行
   divide= func(a,b float64) (float64, error) {
         if b ==0.0{
         	return 0.0,fmt.Errorf("不能除以o")
		 }else {
		 	return a/b,nil
		 }
   }
   //函数类型的调用
	f, err := divide(5.0, 3.0)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)
}
