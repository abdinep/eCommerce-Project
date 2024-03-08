package models

import "time"

type Order struct {
	ID             uint
	User           User
	UserID         int `json:"user_id"`
	Product        Product
	ProductID      int `json:"product_id"`
	Address        Address
	AddressID      int    `json:"address_id"`
	Coupon_Code    string `json:"coupon_code"`
	Order_Quantity int
	Order_Price    int
	Order_Payment  string `json:"order_payment"`
	Order_Date     time.Time
	Order_status   string `json:"status"`
}
