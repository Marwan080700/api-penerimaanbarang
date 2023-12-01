package models

import "time"

type ProductCategories struct {
	ID                  int       `json:"id" gorm:"primaryKey;autoIncrement"`
	NameCategoryProduct string    `json:"product_category_name" form:"product_category_name" gorm:"type:varchar(255)"`
	Desc                string    `json:"desc" form:"desc" gorm:"type:varchar(255)"`
    CreatedBy        string    `json:"created_by" gorm:"type:varchar(255)"`
    CreatedAt        time.Time `json:"-`
    UpdatedBy        string    `json:"updated_by" gorm:"type:varchar(255)"`
    UpdatedAt        time.Time `json:"-`
    Status            int       `json:"status" gorm:"default:0"`
}
