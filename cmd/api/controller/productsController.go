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

func Delete(w http.ResponseWriter, r *http.Request) {
	getId := r.URL.Query().Get("id")
	service.Delete(getId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	getId := r.URL.Query().Get("id")
	product := service.EditProduct(getId)
	templateHtml.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	service.Update(w, r)
}
