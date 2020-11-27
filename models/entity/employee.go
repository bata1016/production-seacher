package entity

import (
	"github.com/jinzhu/gorm"
)

// Employee is a model
type Employee struct {
	gorm.Model
	Name         string
	EmployeeCode string
	Email        string
	Password     string
}
