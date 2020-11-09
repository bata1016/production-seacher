package entity

import "time"

// Production is a model
type Production struct {
	ID             unit
	ProductionCode string
	Category       string
	Company        string
	Width          int
	Height         int
	Price          int
	CreatedAt      time.Time
}
