package models

import "time"

type Offer struct {
	ID        uint
	Product   Product
	ProductId int
	OfferName string `json:"OfferName"`
	Amount    float64 `json:"Amount"`
	Created   time.Time
	Expire    time.Time `json:"Expire"`
}
