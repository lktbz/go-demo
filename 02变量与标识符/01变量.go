package main

import (
        "fmt"
)
//变量声明称全部变量
var ass =20
func main() {
   //变量声明方式：
        //单个变量的声明
        var aa int   //初始变量
        var bb=10   //初始变量并赋值 会进行数据类型的推导
        cc:=20   //进化写法 推导数据类型
        dd:="lktbz"
        eee:="小马哥"
        /**
                注意事项
                :=定义变量只能在函数内部使用，所以经常用var定义全局变量
                Go对已经声明但未使用的变量会在编译阶段报错：** not used
                Go中的标识符以字母或者下划线开头，大小写敏感
                Go推荐使用驼峰命名
         */
        fmt.Println(eee)
        fmt.Println(aa)
        fmt.Println(bb)
        fmt.Println(cc)
        fmt.Println(dd)
        fmt.Println(ass)

        //多变量的定义
        var a ,b string   //多变量声明
        var c,d string ="ljk","sz"   //声明并赋值
        var e ,f int =1,2
        tmed,hh :=12,20  //类型推导并赋值
         g,h:=3,4
         println(tmed,hh)
         println(a)
         println(b)
         println(c)
         println(d)
         println(e)
         println(f)
         println(g)
         println(h)
         //集合类型
         var (
                 ed int
                 gg bool
         )
         println(ed)
         println(gg)

         //变量值的互换
        println("原来的c--"+c)
        println("原来的d--"+d)
         c,d=d,c

         println(c)
         println(d)
        x:=100
         fmt.Println("x的值为：",x)
         fmt.Println("x的值为：",&x) //获取x 变量的内存地址
        //temp,_ = m,n //匿名变量：变量值互换，且丢弃变量n
        //println(temp)
        //_是个特殊的变量名，任何赋予它的值都会被丢弃。该变量不占用命名空间，也不会分配内存。

        _,bbb:=34,35  //将值`35`赋予`b`，并同时丢弃`34`：
        println(bbb)
        //println(_)
        //下面是正确的代码示例：

        //file := bbb
        //in, err := os.Open(string(file))
        //out, err := os.Create(string(file)) // err已经在上方定义，此处的 err其实是赋值
        ////但是如果在第二行赋值的变量名全部和第一行一致，则编译不通过：
        //println(in)
        //println(err)
        //println(out)
        //in, err := os.Open(file)
        //in, err := os.Create(file)     // 即 := 必须确保至少有一个变量是用于声明
        //:=只有对已经在同级词法域声明过的变量才和赋值操作语句等价，如果变量是在外部词法域声明的，那么:=将会在当前词法域重新声明一个新的变量

        //多数据分组书写,常量定义
        //const(
        //       i=100
        //       pi=3.1416
        //       prefix="GO_"
        //
        //)
        //多变量定义
        var(
                i int
                pi float64
                prefix string
        )
        println(i)
        println(pi)
        println(prefix)

        //关键字iota
        //关键字iota声明初始值为0，每行递增1
        const(
                aaa=iota
                bbbb=iota
                ccc=iota
        )

        const (
                ggg=aaa
        )
        println("aaa",aaa,"bbbb",bbbb,"ccc",ccc,"ggg",ggg)



}
