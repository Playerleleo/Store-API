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

func Delete(id string) {
	db := dao.ConectionDatabase()
	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	deleteProduct.Exec(id)
	defer db.Close()
}

func EditProduct(id string) model.Product {
	db := dao.ConectionDatabase()

	productDb, err := db.Query("SELECT * FROM products where id=?", id)
	if err != nil {
		panic(err.Error())
	}

	productToUpdate := model.Product{}

	for productDb.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productDb.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity
	}
	defer db.Close()
	return productToUpdate
}
func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := dao.ConectionDatabase()
	update, err := db.Prepare("UPDATE products SET name=?, description=?, price=?, quantity=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	update.Exec(name, description, price, quantity, id)
	defer db.Close()
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
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

		idConvertedToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error in id conversion:", err)
		}

		UpdateProduct(idConvertedToInt, name, description, priceConvertedToFloat, quantityConvertedToInt)
	}
	http.Redirect(w, r, "/", 301)
}
