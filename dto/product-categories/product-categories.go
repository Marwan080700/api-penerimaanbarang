package productcategoriesdto

type ProductCategoriesRequest struct {
	NameCategoryProduct string `json:"product_category_name" form:"product_category_name" gorm:"type:varchar(255)"`
	Desc                string `json:"desc" form:"desc" gorm:"type:varchar(255)"`
	CreatedBy           string `json:"created_by" gorm:"type:varchar(255)"`
}

type ProductCategoriesResponse struct {
	NameCategoryProduct string `json:"product_category_name" form:"product_category_name" gorm:"type:varchar(255)"`
	Desc                string `json:"desc" form:"desc" gorm:"type:varchar(255)"`
	CreatedBy           string `json:"created_by" gorm:"type:varchar(255)"`
}