package model

import (
	"github.com/bata1016/production-seacher/db"
	"github.com/bata1016/production-seacher/models/entity"
)

// Model Employeeモデルのこと
type Model struct{}

// Employee Employeeの作成
type Employee entity.Employee

// CreateEmployee Employeeのレコード全てを取得
func (m Model) CreateEmployee(resistrationEmployee *Employee) []Employee {
	db := db.GetGormConnect()
	db.Create(&resistrationEmployee)
	defer db.Close()
}
