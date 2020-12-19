package entity

import (
	"github.com/jinzhu/gorm"
)

// Employee is a model
type Employee struct {
	gorm.Model
	Name         string `validate:"required"`
	EmployeeCode string `validate:"required"`
	Email        string `validate:"required,email"`
	Password     string `validate:"required"`
}
