package controller

import (
	"github.com/Store-API/cmd/api/service"
	"html/template"
	"net/http"
)

var templateHtml = template.Must(template.ParseGlob("cmd/api/front/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := service.SelectAllProductsDb()
	templateHtml.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	templateHtml.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	service.CreateProduct(w, r)
}
