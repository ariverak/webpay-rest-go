package sql

import (
	"time"
)

type Product struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255" json:"name"`
	ShortName   string    `gorm:"size:120" json:"shortName"`
	Description string    `json:"description"`
	UnitPrice   int       `json:"unitPrice"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func GetProducts() (products []Product) {
	Db.Find(&products)
	return
}
