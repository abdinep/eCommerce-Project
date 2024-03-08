package models

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	Code      string    `gorm:"unique" json:"code"`
	Discount  float64   `json:"discount"`
	ValidFrom time.Time `json:"validfrom"`
	ValidTo   time.Time `json:"validto"`
}
