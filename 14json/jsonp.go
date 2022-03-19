package main
/**
     官网wiki  有意思，明天接着弄

https://github.com/thedevsaddam/gojsonq/wiki/Queries
 */
import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)
func main() {
	const json = `{"name":{"first":"Tom","last":"Hanks"},"age":61}`
	name:=gojsonq.New().FromString(json).Find("name.first")
	fmt.Println(name)
}
