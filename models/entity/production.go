package entity

import "github.com/jinzhu/gorm"

type Production struct {
	gorm.Model
	ProductionCode string
	Category       string
	Company        string
	Width          int
	Height         int
	Price          int
}
