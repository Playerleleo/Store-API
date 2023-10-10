package service

import (
	"github.com/Store-API/cmd/api/dao"
	"github.com/Store-API/cmd/api/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

func SelectAllProductsDb() []model.Product {
	db := dao.ConectionDatabase()
	selectAllProducts, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}

	p := model.Product{}
	products := []model.Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
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
	defer db.Close()
	return products
}

func InsertTable(name, description string, price float64, quantity int) {
	db := dao.ConectionDatabase()

	insertDatabase, err := db.Prepare("INSERT INTO products (name, description, price, quantity) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insertDatabase.Exec(name, description, price, quantity)
	defer db.Close()
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConvertedToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error in price conversion:", err)
		}

		quantityConvertedToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error in quantity conversion:", err)
		}
		InsertTable(name, description, priceConvertedToFloat, quantityConvertedToInt)
	}
	http.Redirect(w, r, "/", 301)
}
