package customerdto

type CustomerRequest struct {
	IdentityCustomer string `json:"customer_identity" form:"customer_identity" gorm:"type:varchar(255)"`
	NameCustomer     string `json:"customer_name" form:"customer_name" gorm:"type:varchar(255)"`
	EmailCustomer    string `json:"customer_email" form:"customer_email" gorm:"type:varchar(255)"`
	HandponeCustomer string `json:"customer_handpone" form:"customer_handpone" gorm:"type:varchar(255)"`
	NpwpCustomer     string `json:"customer_npwp" form:"customer_npwp" gorm:"type:varchar(255)"`
	AddressCustomer  string `json:"customer_address" form:"customer_address" gorm:"type:varchar(255)"`
	CreatedBy        string `json:"created_by" gorm:"type:varchar(255)" form:"created_by"`
	UpdatedBy        string `json:"updated_by" gorm:"type:varchar(255)" form:"updated_by"`
}

type CustomerResponse struct {
	IdentityCustomer string `json:"customer_identity" form:"customer_identity" gorm:"type:varchar(255)"`
	NameCustomer     string `json:"customer_name" form:"customer_name" gorm:"type:varchar(255)"`
	EmailCustomer    string `json:"customer_email" form:"customer_email" gorm:"type:varchar(255)"`
	HandponeCustomer string `json:"customer_handpone" form:"customer_handpone" gorm:"type:varchar(255)"`
	NpwpCustomer     string `json:"customer_npwp" form:"customer_npwp" gorm:"type:varchar(255)"`
	AddressCustomer  string `json:"customer_address" form:"customer_address" gorm:"type:varchar(255)"`
	CreatedBy        string `json:"created_by" gorm:"type:varchar(255)" form:"created_by"`
	UpdatedBy        string `json:"updated_by" gorm:"type:varchar(255)" form:"updated_by"`
}