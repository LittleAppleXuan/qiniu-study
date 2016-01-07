package db

import (
	"database/sql"

)

func FindByPone(p string) string {
	var num string
	db, err := sql.Open("mysql", "root:12345@tcp(10.3.93.203:3306)/cuc3db?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	stmtOut, err := db.Prepare("select uid from user where phone=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	err = stmtOut.QueryRow(p).Scan(&num)
	if err != nil {
		panic(err.Error())
	}
	return num
}
