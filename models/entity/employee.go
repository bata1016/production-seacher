package entity

import (
	"github.com/jinzhu/gorm"
)

// Employee is a model
type Employee struct {
	gorm.Model
	// ID           uint
	Name         string
	EmployeeCode string
	Email        string
	// CreatedAt    time.Time
}
