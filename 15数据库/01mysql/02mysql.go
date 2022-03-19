package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	. "github.com/lib/pq"
	_ "log"
)
var (
	userName  string = "root"
	password  string = "1234"
	ipAddrees string = "127.0.0.1"
	port      int    = 3306
	dbName    string = "redis"
	charset   string = "utf8"
)
var Db *sqlx.DB
type user struct {
	id int `db:"id"`
	name string`db:"name"`
	age int`db:"age"`
}

func main() {
	//dsn := "root:1234@tcp(127.0.0.1:3306)/redis"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
 	//demo01(Db)
	//execDemo(Db)
	//QueryDemo(Db)
	//QueryReflectDemo(Db)
	//QueryGetAndSelectDemo(Db)
	transactionDemo(Db)
}
/**
   使用sqlx 的第一个例子
 */
func demo01(Db *sqlx.DB)  {
	//单表查询
	sqlStr:="select id,name,age from user where id=?"
	query, err := Db.Query(sqlStr, 1)
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}
	for query.Next() {
		var id ,age int
		var name string
		err := query.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println("get data failed, error:[%v]", err.Error())
		}
		fmt.Println(id,"-",name,"--",age)
	}
}

func execDemo(db *sqlx.DB)  {
	//创建表 的语句
	schema:=` CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer);`
	exec, err := db.Exec(schema)
	if err !=nil {
		fmt.Println("创建失败")
	}
	fmt.Println(exec.RowsAffected())

	cityState := `INSERT INTO place (country, telcode) VALUES (?, ?)`
	countryCity := `INSERT INTO place (country, city, telcode) VALUES (?, ?, ?)`
	//MustExec没有异常处理
	mustExec := db.MustExec(cityState, "Hong Kong", 852)
	fmt.Println(mustExec.LastInsertId())
	db.MustExec(cityState, "Singapore", 65)
	db.MustExec(countryCity, "South Africa", "Johannesburg", 27)
}
/**
  查询demo
 */
func QueryDemo(db *sqlx.DB)  {
	//fetch all places from the db
	rows, err := db.Query("SELECT country, city, telcode FROM place")
	if err != nil {
		fmt.Println("查询失败")
	}
	// iterate over each row
	for rows.Next(){
		//定义接受数据的字段
		var country string
		var city sql.NullString
		var telcode int

		err = rows.Scan(&country, &city, &telcode)
		fmt.Println("查询到对应的数据为：",country,city,telcode)
	}
	err=rows.Err()
}
//定义结构体，与数据库字段进行对应
type Place struct {
	Country       string `db:"country"`
	City           sql.NullString `db:"city"`
	TelephoneCode int `db:"telcode"`
}

//使用反射的方式
func QueryReflectDemo(db *sqlx.DB)  {
	rows, _ := db.Queryx("SELECT * FROM place")
	for rows.Next(){
		var p Place
		rows.StructScan(&p)
		fmt.Printf("获取到的国家为:%s 城市为：%s 电话编码：%d ",p.Country,p.City,p.TelephoneCode)
	}
}

//Get and Select use rows.Scan on scannable types and rows.StructScan on non-scannable types.
//	They are roughly analagous to QueryRow and Query,
//where Get is useful for fetching a single result and scanning it,
//and Select is useful for fetching a slice of results:
//通过上面的官方例子，可以知道，get 获取单一数据，select 获取多条数据
func QueryGetAndSelectDemo(db *sqlx.DB)  {
	p := Place{}
	pp := []Place{}
	err:=db.Get(&p,"select *from place LIMIT 1")
	if err !=nil {
		fmt.Println("查询失败")
	}
	fmt.Println("p 查询出的数据为：",p)
	err = db.Select(&pp, "SELECT * FROM place WHERE telcode > ?", 50)
	fmt.Println("pp查询出的数据为：",pp)
}
// 插入数据
func insertRowDemo(db *sqlx.DB) {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "沙河小王子", 19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}
// 更新数据
func updateRowDemo(db *sqlx.DB) {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 6)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo(db *sqlx.DB) {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

/**
  单表分布式事务，跨库的分布式事务，参见分布式组件seata
  事务操作
*/
func  transactionDemo(db *sqlx.DB) {
	tx, err := db.Beginx() //开启事务
	if err !=nil {
		if tx!=nil{
			//回滚事务
			tx.Rollback()
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=10 where id=?"
	tx.MustExec(sqlStr1, 2)
	sqlStr2 := "Update user set age=200 where id=?"
	tx.MustExec(sqlStr2, 4)
	err = tx.Commit() // 提交事务
	if err!=nil{
		tx.Rollback() // 回滚
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
	fmt.Println("exec trans success!")
}