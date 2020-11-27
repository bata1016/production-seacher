package entity

import "github.com/jinzhu/gorm"

type Production struct {
	gorm.Model
	ProductionCode string
	Category       string
	Company        string
	price          int
	height         int
	width          int
}
