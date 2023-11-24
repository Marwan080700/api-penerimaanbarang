package salesdetaildto

type SalesDetailRequest struct {
	IDSales   int    `json:"sale_id" form:"sale_id" gorm:"type:int"`
	IDProduct int    `json:"product_id" form:"product_id" gorm:"type:int"`
	Qty       int    `json:"qty" form:"qty" gorm:"type:int"`
	Price     int    `json:"price" form:"price" gorm:"type:int"`
	Amount    int    `json:"amount" form:"amount" gorm:"type:int"`
	Desc      string `json:"desc" gorm:"type:varchar(255)"`
	Status    string `json:"status" form:"status" gorm:"type:varchar(255)"`
}

type SalesDetailResponse struct {
	IDSales   int    `json:"sale_id" form:"sale_id" gorm:"type:int"`
	IDProduct int    `json:"product_id" form:"product_id" gorm:"type:int"`
	Qty       int    `json:"qty" form:"qty" gorm:"type:int"`
	Price     int    `json:"price" form:"price" gorm:"type:int"`
	Amount    int    `json:"amount" form:"amount" gorm:"type:int"`
	Desc      string `json:"desc" form:"desc" gorm:"type:varchar(255)"`
	Status    string `json:"status" form:"status" gorm:"type:varchar(255)"`
}