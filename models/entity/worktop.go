package entity

import "github.com/jinzhu/gorm"

type WorkTop struct {
	gorm.Model
	ProductionID   uint
	ProductionCode string
	Category       string
	Company        string
	Width          int
	Price          int
}
