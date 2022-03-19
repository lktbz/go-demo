package main

import "fmt"

type Pperson struct {
	ID string
}
func (p *Pperson)selectorName()string{
	number:=p.ID[:1]
	var sector string
	switch number {
	case "3":
		sector="dev"
	case "5":
		sector="ops"
	default:
		sector="assx"
	}
	return fmt.Sprintf("person id=%v work in %v!",p.ID,sector)
}

func main()  {
p:=	&Pperson{"3346"}
	messahe:=p.selectorName()
	fmt.Println(messahe)
	fmt.Println(p)
}