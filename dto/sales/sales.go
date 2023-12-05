package salesdto

type SalesRequest struct {
	DeliveryOrderNumber string    `json:"delivery_order_number" form:"delivery_order_number" gorm:"type:varchar(255)"`
	IDCustomer          int    `json:"customer_id" form:"customer_id"`
	IDUser              int    `json:"user_id" form:"user_id"`
	DateSale            string `json:"sale_date" form:"sale_date" gorm:"type:varchar(255)"`
	DescriptionSale     string `json:"sale_description" form:"sale_description" gorm:"type:varchar(255)"`
	StatusSale          string `json:"sale_status" form:"sale_status" gorm:"type:varchar(255)"`
	AmountTotal         int    `json:"total_amount" form:"total_amount" gorm:"type:int"`
}

type SalesResponse struct {
	DeliveryOrderNumber string    `json:"delivery_order_number" form:"delivery_order_number" gorm:"type:varchar(255)"`
	IDCustomer          int    `json:"customer_id" form:"customer_id"`
	IDUser              int    `json:"user_id" form:"user_id"`
	DateSale            string `json:"sale_date" form:"sale_date" gorm:"type:varchar(255)"`
	DescriptionSale     string `json:"sale_description" form:"sale_description" gorm:"type:varchar(255)"`
	StatusSale          string `json:"sale_status" form:"sale_status" gorm:"type:varchar(255)"`
	AmountTotal         int    `json:"total_amount" form:"total_amount" gorm:"type:int"`
}