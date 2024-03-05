package models

import "gorm.io/gorm"

type Cart struct{
	gorm.Model
	Product Product
	User User
	Product_Id int `json:"product_id"`
	User_id int `json:"user_id"`
	Quantity int `json:"quantity"`
	Subtotal int `json:"subtotal"`
}