package models

import "time"

type Order struct {
	ID           uint
	User         User
	UserID       int
	Address      Address
	AddressID    int    `json:"address_id"`
	CouponCode   string `json:"coupon_code"`
	OrderPrice   int
	OrderPayment string `json:"order_payment"`
	OrderDate    time.Time
	UpdateDate   time.Time
}
