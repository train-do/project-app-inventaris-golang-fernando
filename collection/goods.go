package collection

import "time"

type Goods struct {
	Id             int
	CategoryId     int
	Name           string
	PhotoUrl       string
	Price          int
	PurchaseDate   time.Time
	TotalUsageDays int
}

var RateDepreciation int = 10
