package entity

import "github.com/jinzhu/gorm"

type Tap struct {
	gorm.Model
	ProductionID   uint
	ProductionCode string
	Category       string
	Company        string
}
