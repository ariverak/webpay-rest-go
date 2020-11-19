package main

import (
	"fmt"
	"time"
	sql "webpay-rest-go/models"

	"gorm.io/datatypes"
)

func main() {
	createBrands()
	createProducts()
}

func createBrands() {
	brands := [3]string{"Nike", "Adidas", "Converse"}
	for i := 0; i < len(brands); i++ {
		sql.CreateBrand(sql.Brand{
			Name: brands[i],
		})
	}
}

func createProducts() {
	converseImages := []string{
		"https://dafitistaticcl-a.akamaihd.net/p/converse-6341-275579-1-product.jpg",
		"https://dafitistaticcl-a.akamaihd.net/p/converse-6343-275579-2-product.jpg",
	}

	expire, _ := time.Parse("2006-01-02", "2020-10-10")

	oneProduct := sql.Product{
		Sku:                "ZACON2",
		Name:               "Zapatilla Urbana Ct As Core Ox Azul Converse",
		Price:              24990,
		SpecialPrice:       19990,
		Description:        "Las mejores tillas",
		BrandID:            1,
		Stock:              20,
		SpecialPriceToDate: expire.UTC(),
		IsFreeShipping:     1,
		Images:             datatypes.JSON([]byte(fmt.Sprintf(`{"images": ["%s","%s"]}`, converseImages[0], converseImages[0]))),
		Details: []sql.ProductDetail{
			sql.ProductDetail{
				Title: "Color Azul",
				Key:   "color",
				Value: "blue",
			},
			sql.ProductDetail{
				Title: "Tamaño",
				Key:   "size",
				Value: "small",
			},
		},
	}
	products := []sql.Product{
		oneProduct,
	}
	for i := 0; i < len(products); i++ {
		sql.CreateProduct(products[i])
	}
}
