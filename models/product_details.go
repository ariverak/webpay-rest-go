package sql

import (
	"time"
)

type ProductDetail struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"size:255" json:"title"`
	Key       string    `gorm:"size:45" json:"key"`
	Value     string    `gorm:"size:255" json:"value"`
	ProductID int       `json:"productId"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
