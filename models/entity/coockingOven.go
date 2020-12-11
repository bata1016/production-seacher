package entity

import "github.com/jinzhu/gorm"

type CookingOven struct {
	gorm.Model
	ProductionID   uint
	ProductionCode string
	Category       string
	Company        string
	width          int
	price          int
}
