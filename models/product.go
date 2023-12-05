package models

import "time"

type Product struct {
	ID                int       `json:"id" gorm:"primaryKey;autoIncrement"`
	IdentityProduct   string    `json:"product_identity" form:"product_identity" gorm:"type:varchar(255)"`
	IDCategoryProduct int       `json:"-" form:"product_category_id"`
	ProductCategories ProductCategories `json:"-" gorm:"foreignKey:IDCategoryProduct"`
	NameProduct       string    `json:"product_name" form:"product_name" gorm:"type:varchar(255)"`
	Unit              string    `json:"unit" form:"unit"  gorm:"type:varchar(255)"`
	Price             int       `json:"price" form:"price" gorm:"type:int"`
	Desc              string    `json:"desc" form:"desc" gorm:"type:varchar(255)"`
    CreatedBy        string    `json:"created_by" gorm:"type:varchar(255)"`
    CreatedAt        time.Time `json:"-`
    UpdatedBy        string    `json:"updated_by" gorm:"type:varchar(255)"`
    UpdatedAt        time.Time `json:"-`
	Status            int       `json:"status" gorm:"default:0"`
}
