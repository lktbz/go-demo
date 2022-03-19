package main

import (
	"os"
	"syscall"
)

/**
  官方定义的 error 接口
 */
type error interface {
	Error() string
}
/**
  路径 错误结构体
 */
//type PathError struct {
//	Op string    // "open", "unlink", etc.
//	Path string  // The associated file.
//	Err error    // Returned by the system call.
//}
/**
  重写接口的方法
 */
//func (e *PathError) Error() string {
//	return e.Op + " " + e.Path + ": " + e.Err.Error()
//}

func main() {
	//筛选出指定的错误类型
 for try:=0;try<2;try++{
	 	_, err := os.Create("zs.txt")
 		if err==nil{
			return
		}
	 if e,ok:=err.(*os.PathError);ok&&e.Err==syscall.ENOSPC {

			continue
	 }
	 return
 }
}
