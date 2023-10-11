package app

import (
	"github.com/Store-API/cmd/api/routs"
	"net/http"
)

func Start() {
	routs.GetroutHome()
	routs.GetroutNewProduct()
	routs.PostProduct()
	routs.DeleteProduct()
	routs.Edit()
	routs.Update()
	http.ListenAndServe(":8080", nil)
}
