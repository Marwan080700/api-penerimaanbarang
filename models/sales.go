package models

import "time"

type Sales struct {
    ID                int       `json:"id" gorm:"primaryKey;autoIncrement"`
    DeliveryOrderNumber int    `json:"delivery_order_number" form:"delivery_order_number" gorm:"type:int"`
    IDCustomer        int       `json:"customer_id" form:"customer_id"`
    Customer          Customer  `json:"-" gorm:"foreignKey:IDCustomer"`
    IDUser            int       `json:"-" form:"user_id"`
    User              User      `json:"-" gorm:"foreignKey:IDUser"`
    DateSale          string    `json:"sale_date" form:"sale_date" gorm:"type:varchar(255)"`
    DescriptionSale   string    `json:"sale_description" form:"sale_description" gorm:"type:varchar(255)"`
    StatusSale        string    `json:"sale_status" form:"sale_status" gorm:"type:varchar(255)"`
    AmountTotal       int       `json:"total_amount" form:"total_amount" gorm:"type:int"`
    CreatedBy         string    `json:"created_by" gorm:"type:varchar(255)"`
    CreatedAt         time.Time `json:"created_at"` // corrected tag
    UpdatedBy         string    `json:"updated_by" gorm:"type:varchar(255)"`
    UpdatedAt         time.Time `json:"updated_at"` // corrected tag
    Status            int       `json:"status" gorm:"default:0"`
}
