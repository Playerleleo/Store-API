package main

import (
	"github.com/Store-API/cmd/api/app"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app.AppRun()
}
