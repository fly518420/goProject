package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func checkErr(err error) {

	if err != nil {
		panic(err)
	}

}

func main() {

	db, err := sql.Open("mysql", "root:123456@/test?charset=utf8")
	checkErr(err)

	// 插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?, departname=?, created=?")
	checkErr(err)

	res, err := stmt.Exec("fly", "研发", "2016-03-16")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	
	// 更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	
	res, err = stmt.Exec("郭鹏飞", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)




	db.Close()
	

}
