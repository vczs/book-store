package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //mysql驱动
)

var (
	Db  *sql.DB
	err error
)

//数据库mysql
func init() {
	Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/book")
	if err != nil {
		panic(err.Error())
	}
}
