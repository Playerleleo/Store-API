package routs

import (
	"github.com/Store-API/cmd/api/controller"
	"net/http"
)

func GetroutHome() {
	http.HandleFunc("/", controller.Index)
}

func GetroutNewProduct() {
	http.HandleFunc("/new", controller.New)
}

func PostProduct() {
	http.HandleFunc("/insert", controller.Insert)
}

func DeleteProduct() {
	http.HandleFunc("/delete", controller.Delete)
}
