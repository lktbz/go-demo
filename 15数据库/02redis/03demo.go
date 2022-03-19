package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	var ctxs = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.5.145:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//查询一个不存在的key 对其进行健壮性检查
	get:=rdb.Get(ctxs, "order")

	fmt.Println(get.Val(),get.Err())

}
