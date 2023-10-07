package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type Product struct {
	Id int
	Name        string
	Description string
	Price      float64
	Quantity    int
}

func conectionDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:admin123@tcp(localhost:3306)/db_store")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success!")
	return db
}

var templateHtml = template.Must(template.ParseGlob("front/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectionDatabase()
	selectAllProducts, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next(){
		var id, quantity int
		var name , description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}



	templateHtml.ExecuteTemplate(w, "Index", products)

	defer db.Close()
}