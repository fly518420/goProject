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

	stmt, err := db.Prepare("insert userinfo set username=?, departname=?, created=?")
	checkErr(err)	
	
	result, err := stmt.Exec("哈哈", "设计", "2016-09-21")
	checkErr(err)

	id, err := result.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	db.Close()
}
