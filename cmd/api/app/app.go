package app

import (
	"github.com/Store-API/cmd/api/routs"
	"net/http"
)

func AppRun() {
	routs.GetroutHome()
	routs.GetroutNewProduct()
	http.ListenAndServe(":8080", nil)
}
