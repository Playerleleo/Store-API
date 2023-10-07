package main

import (
	"html/template"
	"net/http"
)


var templateHtml = template.Must(template.ParseGlob("front/*.html"))

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080",nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
templateHtml.ExecuteTemplate(w, "Index", nil)
}