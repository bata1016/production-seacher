package model

import (
	"github.com/bata1016/production-seacher/db"
	"github.com/bata1016/production-seacher/models/entity"
)

// ここではモデルとデータベースのやりとりを記述

// Model Employeeモデルのこと
type EmployeeModel struct{}

// Employee Employeeの作成
type Employee entity.Employee

// GetAll Employeeを全て取得
func (m EmployeeModel) GetAll() ([]Employee, error) {
	db := db.GetGormConnect()
	var employee []Employee

	if err := db.Find(&employee).Error; err != nil {
		return nil, err
	}

	return employee, nil
}

// CreateModel Employeeを新しく作成
func (m EmployeeModel) CreateModel(name string, employeeCode string, email string, password string) {
	db := db.GetGormConnect()

	err := db.Create(&Employee{Name: name, EmployeeCode: employeeCode, Email: email, Password: password})
	if err != nil {
		panic(err)
	} else {
		defer db.Close()
	}
}
