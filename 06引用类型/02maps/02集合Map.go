package main

import (
        "fmt"
        _ "sync"
)

func main() {
        //Go内置了map类型，map是一个无序键值对集合（也有一些书籍翻译为字典）。

        //创建方式一：
        //* 声明变量，默认 map 是 nil */
        //var map_variable map[key_data_type]value_data_type
        mss:=map[string]int{"a":1,"b":2}
        fmt.Println(mss)
        rating:=map[string]float64 {"C":5, "Go":4.5, "Python":4.5, "C++":2 }
        fmt.Println(rating)


        //make 方式创建
        smake:=make(map[string]string)
        smake["zs"]="zs"
        smake["it"]="beijing"
        smake["jps"]="jpan"
        fmt.Println(smake)
        //遍历
        for key, value := range smake {
           fmt.Printf("key:=%s,value :=%s",key,value)
        }
        //对象方式
        var ms =make(map[string]Peson) //声明
        ms["123"]=Peson{"123","tom"} //赋值操作
        //p,isFind:=ms["123"]
        p:=ms["123"]
        //fmt.Println(isFind)
        fmt.Println(p)
        //注意：golang中map的 key 通常 key 为 int 、string，但也可以是其他类型如：bool、数字、string、
        //指针、channel，还可以是只包含前面几个类型的接口、
        //结构体、数组。slice、map、function由于不能使用 == 来判断，不能作为map的key。

        var countryCapitalMap map[string]string
        /* 创建集合 */
        countryCapitalMap = make(map[string]string)

        /* map 插入 key-value 对，各个国家对应的首都 */
        countryCapitalMap["France"] = "Paris"
        countryCapitalMap["Italy"] = "Rome"
        countryCapitalMap["Japan"] = "Tokyo"
        countryCapitalMap["India"] = "New Delhi"

        //判断某个元素是否存在
        captial, ok:=countryCapitalMap["Japan"]
        /* 如果 ok 是 true, 则存在，否则不存在 */
        if(ok){
                fmt.Println("Capital of United States is", captial)
        }else {
                fmt.Println("Capital of United States is not present")
        }
        /* 创建 map */
        countryCapitalMap1 := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}

        fmt.Println("原始 map")

        /* 打印 map */
        for country := range countryCapitalMap1 {
                fmt.Println("Capital of",country,"is",countryCapitalMap1[country])
        }

        /* 删除元素 */
        delete(countryCapitalMap1,"France");
        fmt.Println("Entry for France is deleted")

        fmt.Println("删除元素后 map")

        /* 打印 map */
        for country := range countryCapitalMap1 {
                fmt.Println("Capital of",country,"is",countryCapitalMap1[country])
        }

        //map 是引用类型
        personSalary := map[string]int{
                "steve": 12000,
                "jamie": 15000,
        }
        personSalary["mike"] = 9000
        fmt.Println("Original person salary", personSalary)
        newPersonSalary := personSalary
        newPersonSalary["mike"] = 18000
        fmt.Println("Person salary changed", personSalary)

       //map 的常用操作
      var numbers map[string]int
       numbers=make(map[string]int)
       numbers["one"]=1
       numbers["two"]=2
       numbers["three"]=3
       delete(numbers,"ten")

       fmt.Println("第三个数字为",numbers["three"])
        /**
        map的遍历：同数组一样，使用for-range 的结构遍历
        注意：
        map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取；
        map的长度是不固定的，也就是和slice一样，也是一种引用类型
        内置的len函数同样适用于map，返回map拥有的key的数量
        go没有提供清空元素的方法，可以重新make一个新的map，不用担心垃圾回收的效率，因为go中并行垃圾回收效率比写一个清空函数高效很多
        map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
         */
        //
        for k, v := range numbers {
	fmt.Println("k的值为=",k," v的值为",v)
        }

        /**
        Go内置的map只有读是线程安全的，读写是线程不安全的。

        需要并发读写时，一般都是加锁，但是这样做性能不高，在go1.9版本中提供了更高效并发安全的sync.Map。

        sync.Map的特点：

        无须初始化，直接声明即可
        sync.Map不能使用map的方式进行取值和设值操作，而是使用sync.Map的方法进行调用。Store表示存储，Load表示获取，Delete表示删除。
        使用Range配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，需要继续迭代时，返回true，终止迭代返回false。
         */

        //var sce sync.Map
        //sce.Store("id","ID162334")
        //sce.Store("name","lktbz")
        //
        ////根据相应的值取值
        //fmt.Println(sce.Load("name"))
        //
        ////遍历操作
        //sce.Range(func(k, v interface{}) bool {
	//fmt.Println(" sync.Map k的值为=",k," v的值为",v)
	//return true
        //})








}

type Peson struct {
        ID string
        Name string
}


