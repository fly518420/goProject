package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/test?charset-utf8")
	checkErr(err)

	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	_, err = stmt.Exec("飞fly", 1)
	checkErr(err)

	db.Close()
}
