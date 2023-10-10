package routs

import (
	"github.com/Store-API/controller"
	"net/http"
)

func GetroutHome() {
	http.HandleFunc("/", controller.Index)
}

func GetroutNewProduct() {
	http.HandleFunc("/new", controller.New)
}
