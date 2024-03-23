package models

import "time"

type Offer struct {
	ID         uint
	OfferType  string
	ProductId  int
	CategoryID int
	Amount     float64
	Created    time.Time
	Expire     time.Time
}
