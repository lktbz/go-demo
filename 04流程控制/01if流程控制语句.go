package main

import "fmt"

func main() {

        //三种if
        /**
          if condition{

           }
        */
        /**
        if condition {
            // do something
        } else {
            // do something
        }
        */

        /**
          	if condition1 {
                // do something
            } else if condition2 {
                // do something else
            }else {
                // catch-all or default
            }
        */

        //var i int =4
        var i int = 3
        if i == 3 {
	println("i 的值为3")
        } else {
	println("i 的值非3")
        }

        //特殊写法

        //if 还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断，代码如下：
        //if err := Connect(); err != nil {
        //	fmt.Println(err)
        //	return
        //}
        //Connect 是一个带有返回值的函数，err:=Connect() 是一个语句，执行 Connect 后，将错误保存到 err 变量中。
        //err != nil 才是 if 的判断表达式，当 err 不为空时，打印错误并返回。
        //
        if nu:=4;nu%2==0{
	fmt.Println(nu,"is even")
        }else{
                fmt.Println(nu," is odd")
        }


        //这种写法可以将返回值与判断放在一行进行处理，而且返回值的作用范围被限制在 if、else 语句组合中。
        nums := 4
        switch nums {
        case 1:
	fmt.Print("lala")
        case 2:
	fmt.Print("lblb")
        case 3:
	fmt.Print("lclc")
        default:
	fmt.Print("ldld")

        }
        //Go保留了break，用来跳出switch语句，上述案例的分支中默认就书写了该关键字
        //Go也提供fallthrough，代表不跳出switch，后面的语句无条件执行
        var grade string = "B"
        switch  {
        case grade=="A":
	fmt.Println("优秀\n")
        case grade=="B",grade=="C":
                fmt.Println("良好")
        case grade=="D":
                fmt.Println("及格")
        case grade=="E":
                fmt.Println("不及格")
        default:
                fmt.Println("差")
        }
        fmt.Println("你的等级是",grade)



	  if n1,n2:="d","a";n1==n2{
	  	fmt.Println("n1 and n2 相等")
	  }else if n1=="d"&& n2=="a"{
	  	fmt.Println("n1==d,n2==a")
	  }else{
	  	fmt.Println("n1,n2 既不是a也不是b")
	  }

	  //高级用法
	  //初始化条件变量                 检查条件变量是否符合条件
	  if n3,n4,check:=413,404,true;check{
	  	fmt.Println(n3+n4) //符合条件执行语句
	  }
	if n3,n4,check:=413,404,false;!check{
		fmt.Println(n3-n4) //符合条件执行语句
	}
}
