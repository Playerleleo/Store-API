package app

import (
	"github.com/Store-API/cmd/api/routs"
	"net/http"
)

func Start() {
	routs.GetroutHome()
	routs.GetroutNewProduct()
	http.ListenAndServe(":8080", nil)
}
