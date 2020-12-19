package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/jinzhu/gorm"
)

type Production struct {
	gorm.Model
	ProductionCode string `gorm:"not null"`
	Category       string `gorm:"not null"`
	Company        string `gorm:"not null"`
	Width          string
	Height         string
	Price          string `gorm:"not null"`
}

func (p Production) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.ProductionCode, validation.Required),
		validation.Field(&p.Category, validation.Required),
		validation.Field(&p.Width, is.Int),
		validation.Field(&p.Height, is.Int),
		validation.Field(&p.Price, validation.Required, is.Int),
	)
}
