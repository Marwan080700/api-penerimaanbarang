package productdto

type ProductRequest struct {
	IdentityProduct   string `json:"product_identity" gorm:"type:varchar(255)" form:"product_identity"`
	IDCategoryProduct int    `json:"product_category_id" gorm:"type:varchar(255)" form:"product_category_id"`
	NameProduct       string `json:"product_name" gorm:"type:varchar(255)" form:"product_name"`
	Unit              string    `json:"unit" gorm:"type:varchar(255)" form:"unit"`
	Price             int    `json:"price" gorm:"type:int" form:"price"`
	Desc              string `json:"desc" gorm:"type:varchar(255)" form:"desc"`
	CreatedBy         string `json:"created_by" gorm:"type:varchar(255)" form:"created_by"`
	UpdatedBy         string `json:"updated_by" gorm:"type:varchar(255)" form:"updated_by"`
}

type ProductResponse struct {
	IdentityProduct   string `json:"product_identity" gorm:"type:varchar(255)" form:"product_identity"`
	IDCategoryProduct int    `json:"product_category_id" gorm:"type:varchar(255)" form:"product_category_id"`
	NameProduct       string `json:"product_name" gorm:"type:varchar(255)" form:"product_name"`
	Unit              string    `json:"unit" gorm:"type:varchar(255)" form:"unit"`
	Price             int    `json:"price" gorm:"type:int" form:"price"`
	Desc              string `json:"desc" gorm:"type:varchar(255)" form:"desc"`
	CreatedBy         string `json:"created_by" gorm:"type:varchar(255)" form:"created_by"`
	UpdatedBy         string `json:"updated_by" gorm:"type:varchar(255)" form:"updated_by"`
}
