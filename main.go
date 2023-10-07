package main

import (
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price      float64
	Quantity    int
}

var templateHtml = template.Must(template.ParseGlob("front/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
    produtos := []Product{
        {Name: "Camiseta", Description: "Azul, bem bonita", Price: 39, Quantity: 5},
    }

    templateHtml.ExecuteTemplate(w, "Index", produtos)
}