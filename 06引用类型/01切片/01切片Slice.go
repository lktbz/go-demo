package main

import "fmt"

func main() {
        //切片(slice)解决了数组长度不能扩展，以及基本类型数组传递时产生副本的问题
   //定义方式
        var s1[]int // 和声明数组一样，只是没有长度，但是这样做没有意义，因为底层的数组指针为nil
        s2:=[]byte{'a','b','c'}
        fmt.Println(s1)
        fmt.Println(s2)

        //使用make 函数创建
        slice1 := make([]int,5)		// 创建长度为5，容量为5，初始值为0的切片
        slice2 := make([]int,5,7)	// 创建长度为5，容量为7，初始值为0的切片
        slice3 := []int{1,2,3,4,5}	// 创建长度为5，容量为5，并已经初始化的切片
        fmt.Println(slice1)
        fmt.Println(slice2)
        fmt.Println(slice3)

        slice4s:=make([]int,5)
        fmt.Println("slice4的值为:",slice4s)
        slice4s=append(slice4s,1,2,3,4,5)
        for _, slice4 := range slice4s {
             fmt.Println("for 遍历的值为：",slice4)
        }


        //从数组创建：slice可以从一个数组再次声明。slice通过array[i:j]来获取，
        //其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i:
        var arr = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
        var a,b,c,d[]byte  //声明两个含有byte的slice
        a=arr[2:5] //a指向数组的第3个元素开始，并到第五个元素结束，现在a含有的元素: ar[2]、ar[3]和ar[4]
        b=arr[3:5]//b是数组arr的另一个slicre,b的元素是：ar[3]和ar[4]
        //[:]冒号前面表示切片中有从第几个元素开始，后面表示元素的长度
        c=arr[:7]
        d=arr[5:]
        fmt.Println(string(a))
        fmt.Println(string(b))
        fmt.Println(string(c))
        fmt.Println(string(d))


        slice4 := make([]int, 5, 10)
        fmt.Println(len(slice4))			// 5
        fmt.Println(cap(slice4))			// 10文件
        fmt.Println(slice4)

        //增加
        slice1= append(slice1,1,2)
        fmt.Println(slice1)

        //增加一个新的切片
        slice2tmp:=make([]int,3)
        //将新建的切片放进原来创建的切片中
        slice1=append(slice1,slice2tmp...)
        fmt.Println(slice1)

        //切片的拷贝
        s4:=[]int{1,2,3}
        s3:=make([]int,3)
        nums:=copy(s3,s4)
        fmt.Println(s3)
        fmt.Println(s4)
        fmt.Println(nums)
        // 声明一个数组
       // var array = [10文件]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
        // 声明两个slice
        //var aSlice, bcs_ []byte

        // 演示一些简便操作
       // aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
       // aSlice = array[5:] // 等价于aSlice = array[5:10文件] aSlice包含元素: f,g,h,i,j
        //aSlice = array[:] // 等价于aSlice = array[0:10文件] 这样aSlice包含了全部的元素

        // 从slice中获取slice
        //aSlice = array[3:7] // aSlice包含元素: d,e,f,g，len=4，cap=7
        //bcs__ = aSlice[1:3]     // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
        //bcs__ = aSlice[:3]      // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
        //bcs__ = aSlice[0:5]     // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
        //bcs__ = aSlice[:]       // bSlice包含所有aSlice的元素: d,e,f,g
}


