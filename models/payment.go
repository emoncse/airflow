package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	ProductName string `json:"product_name"`
	Amount      int    `json:"amount"`
}
