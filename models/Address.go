package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Address string `json:"user_address"`
	City    string `json:"user_city"`
	State   string `json:"user_state"`
	Pincode int    `json:"user_pincode"`
	Country string `json:"user_country"`
	Phone   int    `json:"user_phone"`
	UserId  int    `json:"user_id"`
}
