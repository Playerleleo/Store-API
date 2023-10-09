package service

import (
	"github.com/Store-API/dao"
	"github.com/Store-API/model"
	_ "github.com/go-sql-driver/mysql"
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
