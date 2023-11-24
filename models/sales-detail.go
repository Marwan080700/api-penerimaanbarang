package models

import "time"

type SalesDetail struct {
    ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
    Sales     Sales     `json:"-" gorm:"foreignKey:IDSales"`
    IDSales   int       `json:"sale_id" form:"sale_id" gorm:"type:int"`
    Product   Product   `json:"-" gorm:"foreignKey:IDProduct"`
    IDProduct int       `json:"-" form:"product_id" gorm:"type:int"`
    Qty       int       `json:"qty" form:"qty" gorm:"type:int"`
    Price     int       `json:"price" form:"price" gorm:"type:int"`
    Amount    int       `json:"amount" form:"amount" gorm:"type:int"`
    Desc      string    `json:"desc" form:"desc" gorm:"type:varchar(255)"`
    Status    string    `json:"status" form:"status" gorm:"type:varchar(255)"`
    CreatedBy string    `json:"created_by" gorm:"type:varchar(255)"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedBy string    `json:"updated_by" gorm:"type:varchar(255)"`
    UpdatedAt time.Time `json:"updated_at"`
}
