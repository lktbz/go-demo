package main

import (
	"fmt"
	"net/http"
)

func main() {
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello go!"))
})
	err := http.ListenAndServe(":8080", nil)
	if err!=nil{
		panic(err.Error()) //直接抛错误，并终止程序运行
	}
	fmt.Println("go is started")
}
//当起连个服务的时候，就发生了错误
