package models

import "time"

type Order struct {
	OrderID      int
	CustomerName string
	Item         []Item
	OrderedAt    time.Time
}
