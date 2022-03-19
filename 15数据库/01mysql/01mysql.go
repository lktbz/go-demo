package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
/**
简单的查询操作
 */
func main() {
	db, err := sql.Open("mysql", "root:1234@/redis?charset=utf8")
	if err != nil {
		 fmt.Println(err)
	}
	sqlstr:="insert into user(name,age)values (?,?)"
	exec, err := db.Exec(sqlstr, "王五", 25, )
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	id, err := exec.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", id)
}
