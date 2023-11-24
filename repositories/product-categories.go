package repositories

import (
	"pengirimanbarang/models"

	"gorm.io/gorm"
)

type ProductCategoriesRepository interface {
	FindProductCategories() ([]models.ProductCategories, error)
	GetProductCategories(ID int) (models.ProductCategories, error)
	CreateProductCategories(productcategories models.ProductCategories) (models.ProductCategories, error)
	UpdateProductCategories(productcategories models.ProductCategories) (models.ProductCategories, error)
	DeleteProductCategories(productcategories models.ProductCategories) (models.ProductCategories, error)
	DeleteProductsByCategoryID(categoryID int) error
}

func RepositoryProductCategories(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProductCategories() ([]models.ProductCategories, error) {
	var productcategories []models.ProductCategories
	err := r.db.Find(&productcategories).Error

	return productcategories, err
}

func (r *repository) GetProductCategories(ID int) (models.ProductCategories, error) {
	var productcategories models.ProductCategories
	err := r.db.First(&productcategories, ID).Error

	return productcategories, err
}

func (r *repository) CreateProductCategories(productcategories models.ProductCategories) (models.ProductCategories, error) {
    err := r.db.Create(&productcategories).Error

    return productcategories, err
}

func (r *repository) DeleteProductCategories(productcategories models.ProductCategories) (models.ProductCategories, error) {
	err := r.db.Delete(&productcategories).Error

	return productcategories, err
}

func (r *repository) UpdateProductCategories(productcategories models.ProductCategories) (models.ProductCategories, error) {
	err := r.db.Save(&productcategories).Error

	return productcategories, err
}

func (r *repository) DeleteProductsByCategoryID(categoryID int) error {
    // Assuming products have a foreign key to the product category
    err := r.db.Where("id_category_product = ?", categoryID).Delete(&models.Product{}).Error
    return err
}
