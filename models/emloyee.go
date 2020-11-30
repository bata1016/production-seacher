package model

import (
	"crypto"

	"github.com/bata1016/production-seacher/db"
	"github.com/bata1016/production-seacher/models/entity"
)

// ここではモデルとデータベースのやりとりを記述

// EmployeeModel Model Employeeモデルのこと
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
	passwordEncrypt, _ := crypto.PasswordEncrypt(password)
	db.Create(&Employee{Name: name, EmployeeCode: employeeCode, Email: email, Password: passwordEncrypt})
	defer db.Close()
	// defer db.Close()
	// err := db.Create(&Employee{Name: name, EmployeeCode: employeeCode, Email: email, Password: password}).Error
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	defer db.Close()
	// }
}

func (m EmployeeModel) GetLoginEmployee(employeeCode string) Employee {
	db := db.GetGormConnect()
	var employee Employee
	db.First(&employee, "employee_code = ?", employeeCode)
	// db.First(&employee, "password=?", password)
	db.Close()
	return employee
}
