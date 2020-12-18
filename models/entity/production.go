package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type Production struct {
	gorm.Model
	ProductionCode string `gorm:"not null"`
	Category       string `gorm:"not null"`
	Company        string `gorm:"not null"`
	Width          int
	Height         int
	Price          int `gorm:"not null"`
}

func (p Production) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.ProductionCode, validation.Required),
		validation.Field(&p.Category, validation.Required),
		validation.Field(&p.Price, validation.Required),
	)
}
