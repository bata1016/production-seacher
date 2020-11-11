package model

import (
	"github.com/bata1016/production-seacher/db"
	"github.com/bata1016/production-seacher/models/entity"
	"github.com/gin-gonic/gin"
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
func (m EmployeeModel) CreateModel(ctx *gin.Context) (Employee, error) {
	db := db.GetGormConnect()
	var employee Employee

	if err := ctx.BindJSON(&employee); err != nil {
		return employee, nil
	}

	if err := db.Create(&employee).Error; err != nil {
		return employee, err
	}

	return employee, nil
}
