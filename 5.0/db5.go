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

	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string
		err = rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)
		fmt.Printf("uid=%d, username=%s, departname=%s, created=%s", uid, username, departname, created)
		fmt.Println()
	}

	db.Close()
}
