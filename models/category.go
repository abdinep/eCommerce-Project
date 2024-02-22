package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Name        string `gorm:"unique" json:"categoryName"`
	Description string `json:"catDescription"`
	Status      string `gorm:"default:Active" json:"catStatus"`
}
