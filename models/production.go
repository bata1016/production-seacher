package model

import (
	"github.com/bata1016/production-seacher/db"
	"github.com/bata1016/production-seacher/models/entity"
)

// ProductionModel Model Productionのモデル
type ProductionModel struct{}

// Production Productionの作成
type Production entity.Production

func (p ProductionModel) GetAll() ([]Production, error) {
	db := db.GetGormConnect()
	var production []Production

	if err := db.Find(&production).Error; err != nil {
		return nil, err
	} else {
		db.Close()
		return production, err
	}
}

func (p ProductionModel) CreateProductionModel(productionCode string, category string, company string, width string, height string, price string) error {
	db := db.GetGormConnect()
	productionParams := &entity.Production{
		ProductionCode: productionCode, Category: category, Company: company, Width: width, Height: height, Price: price,
	}
	errorMassages := productionParams.Validate()
	if errorMassages != nil {
		return errorMassages
	} else {
		db.Create(&Production{ProductionCode: productionCode, Category: category, Company: company, Width: width, Height: height, Price: price})
		db.Close()
		return errorMassages
	}
}
