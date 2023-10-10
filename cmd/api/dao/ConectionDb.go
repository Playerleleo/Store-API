package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConectionDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:admin123@tcp(localhost:3306)/db_store")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success!")
	return db
}
