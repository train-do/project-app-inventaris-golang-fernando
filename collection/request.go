package collection

import "time"

type FormGoods struct {
	Name         string
	CategoryId   int
	PhotoUrl     string
	Price        int
	PurchaseDate time.Time
}
type FormCategory struct {
	Name        string
	Description string
}
