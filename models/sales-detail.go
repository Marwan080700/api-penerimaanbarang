package models

import "time"

type SalesDetail struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	IDSales   int       `json:"-" form:"sale_id" gorm:"type:int"`
	Sales     Sales		`json:"-" gorm:"foreignKey:IDSales"`
	IDProduct int       `json:"-" form:"product_id" gorm:"type:int"`
	Product   Product	`json:"-" gorm:"foreignKey:IDProduct"`
	Qty       int       `json:"qty" form:"qty" gorm:"type:int"`
	Price     int       `json:"price" form:"price" gorm:"type:int"`
	Amount    int       `json:"amount" form:"amount" gorm:"type:int"`
	Desc      string       `json:"desc" form:"desc" gorm:"type:varchar(255)"`
	Status    string       `json:"status" form:"status" gorm:"type:varchar(255)"`
    CreatedBy        string    `json:"created_by" gorm:"type:varchar(255)"`
    CreatedAt        time.Time `json:"-`
    UpdatedBy        string    `json:"updated_by" gorm:"type:varchar(255)"`
    UpdatedAt        time.Time `json:"-`
}
