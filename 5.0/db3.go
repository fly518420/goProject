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

	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	
	result, err := stmt.Exec(1)
	checkErr(err)

	affect,	err := result.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}
