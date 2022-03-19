package main

import "fmt"

/**
   定义接口
 */
type Writer interface {
    Writer([]byte)(int,error)
}
/**
定义结构体
 */
type ConsoleWriter struct {
}

/**
  定义方法
 */

func (cw ConsoleWriter)Writer(data []byte)(int ,error)  {
 n,err:=	fmt.Println(string(data))
	return n,err
}
func main() {
	sc:=ConsoleWriter{}
	sc.Writer([]byte("hello golang"))
}
