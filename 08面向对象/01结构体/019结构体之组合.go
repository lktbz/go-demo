package main

import (
	"fmt"
	"time"
)

type Building struct {
	ID string
	Districicode int
	Builded time.Time
	Companies[]string
	Price float32
	Apartments map[int]bool
}
func main() {
  b:=&Building{"1545",3,time.Now(),
   	[]string{"unit","lajs"},
  541.00,
	  map[int]bool{
  	   1:true,
  	   2:false,
	  },
  }
  fmt.Println(b)
  fmt.Println(b.Companies)
  fmt.Println(b.Apartments)
}
