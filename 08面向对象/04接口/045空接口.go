package main

type Uss struct {
	Id interface{} //可以是任何数据类型
}
//返回了一个string 类型
func (uss *Uss)sendMessage(interface{}) interface{} {
	return "log"
}

