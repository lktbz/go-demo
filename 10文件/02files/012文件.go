package main

import (
	"fmt"
	"io"
	"os"
	"path"
)

//文件的写入：
func main() {
	//写入字节
	WriteByByte()
	WriteByString()
	//SeakDemo()
	//FileInfo()
	//Dit()
	//deleteDir()
}
func deleteDir(){
	err := os.Remove("test.txt")
	if err!=nil{
		fmt.Println("移除失败:",err)
	}
}
/**
   文件路径目录操作
 */
func Dit()  {
	abs := path.IsAbs("./10文件/path.file") //路径是否绝对路径
	fmt.Println(abs)
  	//创建目录
    os.Mkdir("./testcreatedir",os.ModePerm)
	//多级目录的创建
	err := os.MkdirAll("./level1/level2/level3", os.ModeAppend)
	if err!=nil{
		fmt.Println("目录创建失败")
	}
}
/**
   获取文件的描述信息
 */
func FileInfo()  {
	stat, _ := os.Stat("test.txt")
	fmt.Println("文件的信息包含：",stat)
	fmt.Println("文件名",stat.Name())
	fmt.Println("文件大小",stat.Size())
	fmt.Println("修改时间？",stat.ModTime())
	fmt.Println("文件权限？",stat.Mode())
	fmt.Println("是否目录？",stat.IsDir())
	fmt.Println("系统信息？",stat.Sys())
}
/**
修改文件的读写指针位置 Seek()，包含两个参数：

	参数1：偏移量，为正数时向文件尾偏移，为负数时向文件头偏移
	参数2：偏移的开始位置，包括：
	io.SeekStart：从文件起始位置开始
	io.SeekCurrent：从文件当前位置开始
	io.SeekEnd：从文件末尾位置开始

 */
func SeakDemo(){
	f, _ := os.OpenFile("test.txt",os.O_RDWR, 6)
	off, _ := f.Seek(5, io.SeekStart)
	fmt.Println(off)							// 5
	n, _ := f.WriteAt([]byte("111"), off)
	fmt.Println(n)
	f.Close()
}

/**
   写入字符串
 */

func  WriteByString() {
		//追加可写模式
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_RDWR, 6)
	if err!=nil{
		fmt.Println("打开文件失败",err)
	}
	writeString, err := file.WriteString("\n这是写入字符串测试")
	 if err !=nil{
	 	fmt.Println("文件写入失败",err)
	 }
	fmt.Println("写入的数量为：",writeString)

}
/**
    字节写入
 */
func  WriteByByte()  {
	file, err := os.OpenFile("test.txt", os.O_RDWR, 6)
	if err!=nil{
		fmt.Println("err:",err)
	}else{
		//写入字节
		write, err := file.Write([]byte("张三小伙子"))
		if err!=nil{
			fmt.Println("写入失败：",err)
		}
		fmt.Println("写入的是：",write)
	}
}
