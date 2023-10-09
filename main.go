package main

import (
	"github.com/Store-API/routs"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	routs.Getrouts()
	http.ListenAndServe(":8080", nil)
}
