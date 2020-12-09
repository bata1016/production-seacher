package entity

import "github.com/jinzhu/gorm"

type Production struct {
	gorm.Model
	WorkTops []WorkTop
}
