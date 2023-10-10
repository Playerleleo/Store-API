package main

import (
	"github.com/Store-API/routs"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	routs.GetroutHome()
	routs.GetroutNewProduct()
	http.ListenAndServe(":8080", nil)
}
