package main

import (
	"fmt"
	"time"
)

type Metting struct {
	ID string
	Date time.Time
}
func main() {
  p0:=&Metting{"311411",time.Now()}
  fmt.Println(p0)



  list:=[]*Metting{
  	p0,
  	&Metting{"41121",time.Now()},
  }
  fmt.Println(list)

	for _, metting := range list {
		fmt.Println(metting.Date,metting.ID)
	}


}
