package main

import "fmt"

type  Entity struct {
	Id uint64
	Age uint8

}
type man struct {
	*Entity
	Name string
}
type Anis struct {
	Entity
	Species string
}
func main() {

	ms:=man{&Entity{Id: 1341541,Age: 28},"lk"}
	ans:=Anis{Entity{Id: 14454,Age: 52,
	},"jd"}
	fmt.Println(ms)
	fmt.Println(ans.Age)
	fmt.Println(*ms.Entity)

}
