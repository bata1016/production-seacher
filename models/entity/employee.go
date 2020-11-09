package entity

import "time"

// Employee is a model
type Employee struct {
	ID           uint
	Name         string
	EmployeeCode string
	Email        string
	CreatedAt    time.Time
}
