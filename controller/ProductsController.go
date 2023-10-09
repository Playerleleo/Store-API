package controller

import (
	"github.com/Store-API/service"
	"html/template"
	"net/http"
)

var templateHtml = template.Must(template.ParseGlob("front/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := service.SelectAllProductsDb()
	templateHtml.ExecuteTemplate(w, "Index", allProducts)
}
