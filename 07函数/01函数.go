package main

import (
	"bytes"
	"fmt"
)

/**

  语法格式
func funcName(parametername type1, parametername type2) (output1 type1, output2 type2) {
//这里是处理逻辑代码
//返回多个值
return value1, value2
}
 */
//无返回值，无参数
func fn()  {
	
}
//有返回值无参数
func fn1() (result int) {
	return  1
}
//另一种方式
func fn2()(result int)  {
 	result =2
 	return result
}
//多返回值
//三种不同的写法
func  fn3()(int ,int ,int)  {
 return 1,2,3
}
func fn4()(re int,rf int,rs int)  {
	return  1,2,3
}
func fn5()(re int,rf int,rs int)  {
	re=1
	rf=2
	rs=3
	return  re,rf,rs
}

//可变参数
func joinStrings(Slice...string)string{
  var buf bytes.Buffer
	for _, s := range Slice {
		buf.WriteString(s)
	}
	return buf.String()
}

//参数的传递
// 实际打印函数
func rawPrint(rawList ...interface{}) {
	for _, a := range rawList {
		fmt.Println(a)
	}
}

// 封装打印函数
func print(slist ...interface{})  {
	// 将slist可变参数切片完整传递给下一个函数
	rawPrint(slist...)
}

func main() {
   fmt.Println(joinStrings("pig","-and-","bird"))
	print(1,2,3)
}
