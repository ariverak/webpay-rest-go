package models

import (
	"time"

	"gorm.io/datatypes"
)

type Order struct {
	ID              uint
	SessionID       string
	BuyOrder        string
	FormAction      string
	TokenWs         string
	Status          string
	Discount        uint
	Total           uint
	WebPayData      datatypes.JSON
	BillingAddress  string
	BillingAddress2 string
	BillingAddress3 string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
