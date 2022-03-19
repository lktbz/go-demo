package main

import "fmt"

func main() {
    name:="hello world"
    fmt.Println(name)

    fmt.Println(len(name))

    fmt.Println(name[1])

    a:='h'
    b:=104
    fmt.Printf("%c,%c,%c\n",name[1],a,b)

    //字符串的遍历
    for i := 0; i <len(name) ; i++ {
        fmt.Printf("%c",name[i])
    }
    fmt.Println("------------")
    //for range遍历
    for _, v := range name {
        fmt.Printf("%c",v)
    }
    fmt.Println("------------")
    //字符串是子节的集合？
    slice:=[]byte{65,66,67,68,69}
    s3:=string(slice)
    fmt.Println(s3)


    s4 := "abcdef"
    seke:=[]byte(s4)
    fmt.Println(seke)
}
