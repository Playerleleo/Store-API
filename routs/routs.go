package routs

import (
	"github.com/Store-API/controller"
	"net/http"
)

func Getrouts() {
	http.HandleFunc("/", controller.Index)
}
