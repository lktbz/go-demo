package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}
func main() {
var movies=[]Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}
	marshal, err := json.MarshalIndent(movies,"","  ")
	if err!=nil{
		fmt.Println("解析出错")
	}
	fmt.Println(string(marshal))

   //解析某个字段
	var titles[]struct{Title string}
	 if err:=json.Unmarshal(marshal,&titles);err!=nil{
	 	fmt.Println("err",err)
	 }
	 fmt.Println(titles)

}
