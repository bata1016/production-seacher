package entity

import "github.com/jinzhu/gorm"

type Door struct {
	gorm.Model
	ProductionID   uint
	ProductionCode string
	Category       string
	Company        string
	width          int
	height         int
	price          int
}
