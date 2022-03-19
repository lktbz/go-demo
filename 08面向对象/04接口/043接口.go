package main
/**
   接口嵌套
 */
//定义一个写接口
type Writer interface {
	Write(p []byte) (n int, e error)
}
type Reader interface {
	Read() error
}
//接口的嵌套
type IO interface {
		Writer
		Reader
}
func main() {

}
