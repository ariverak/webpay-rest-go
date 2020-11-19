package sql

import (
	"time"

	"gorm.io/datatypes"
)

type Product struct {
	ID                 int             `gorm:"primaryKey" json:"id"`
	Sku                string          `json:"sku"`
	Name               string          `gorm:"size:255" json:"name"`
	Description        string          `json:"description"`
	BrandID            int             `json:"brandId"`
	Brand              Brand           `gorm:"foreignKey:BrandID" json:"brand"`
	Images             datatypes.JSON  `json:"images"`
	Price              int             `json:"price"`
	Stock              int             `json:"stock"`
	SpecialPrice       int             `json:"specialPrice"`
	SpecialPriceToDate time.Time       `json:"specialPriceToDate"`
	IsFreeShipping     int             `json:"isFreeShipping"`
	Details            []ProductDetail `json:"details"`
	CreatedAt          time.Time       `json:"createdAt"`
	UpdatedAt          time.Time       `json:"updatedAt"`
}

func GetProducts() (products []Product) {
	Db.Find(&products)
	return
}

func CreateProduct(product Product) Product {
	Db.Create(&product)
	return product
}
