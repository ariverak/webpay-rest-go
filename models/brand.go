package sql

import (
	"time"
)

type Brand struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255" json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func GetBrands() (brands []Brand) {
	Db.Find(&brands)
	return
}

func CreateBrand(brand Brand) Brand {
	Db.Create(&brand)
	return brand
}
