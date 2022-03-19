package main

import "fmt"

type Team struct {
	Name string
}
func (t Team)GetName()string{
	return t.Name
}

func (t *Team)SetName(name string){
	t.Name=name
}
func main()  {
  t1:=&Team{}
  t1.SetName("A413")
  fmt.Println(t1.GetName())


  t1.Name="B413"
  name2:=t1.Name
  fmt.Println(name2)




}