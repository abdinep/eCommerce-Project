package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Product_Name string     `gorm:"not null" json:"prodName"`
	Price        int        `json:"price"`
	Quantity     int        `json:"quantity"`
	Size         int        `json:"size"`
	Description  string     `gorm:"not null" json:"description"`
	Category_id  uint        `gorm:"not null" json:"category"`
	Category     Categories 
}
