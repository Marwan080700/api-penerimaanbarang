package models

import "time"

type Customer struct {
	ID               int       `json:"id" gorm:"primaryKey;autoIncrement"`
	IdentityCustomer string    `json:"customer_identity" form:"customer_identity" gorm:"type:varchar(255)"`
	NameCustomer     string    `json:"customer_name" form:"customer_name" gorm:"type:varchar(255)"`
	EmailCustomer    string    `json:"customer_email" form:"customer_email" gorm:"type:varchar(255)"`
	HandponeCustomer string    `json:"customer_handpone" form:"customer_handpone" gorm:"type:varchar(255)"`
	NpwpCustomer     string    `json:"customer_npwp" form:"customer_npwp" gorm:"type:varchar(255)"`
	AddressCustomer  string    `json:"customer_address" form:"customer_address" gorm:"type:varchar(255)"`
    CreatedBy        string    `json:"created_by" gorm:"type:varchar(255)"`
    CreatedAt        time.Time `json:"-`
    UpdatedBy        string    `json:"updated_by" gorm:"type:varchar(255)"`
    UpdatedAt        time.Time `json:"-`
	Status            int       `json:"status" gorm:"default:0"`
}
