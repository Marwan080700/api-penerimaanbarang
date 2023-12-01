package repositories

import (
	"errors"
	"pengirimanbarang/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProducts() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(product models.Product) (models.Product, error)
	GetProductsByCategoryID(categoryID int) ([]models.Product, error)

}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error

	return products, err
}

func (r *repository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, ID).Error

	return product, err
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
    // Check if the product category exists before creating the product
    var categoryCount int64
    if err := r.db.Model(&models.ProductCategories{}).Where("id = ?", product.IDCategoryProduct).Count(&categoryCount).Error; err != nil {
        return models.Product{}, err
    }

    if categoryCount == 0 {
        // The product category doesn't exist, so you should handle this error or return an appropriate response.
        return models.Product{}, errors.New("Product category not found")
    }

    err := r.db.Create(&product).Error

    return product, err
}


func (r *repository) DeleteProduct(product models.Product) (models.Product, error) {
	err := r.db.Delete(&product).Error

	return product, err
}

func (r *repository) UpdateProduct(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error

	return product, err
}

func (r *repository) GetProductsByCategoryID(categoryID int) ([]models.Product, error) {
    var products []models.Product
    err := r.db.Where("id_category_product = ?", categoryID).Find(&products).Error

    return products, err
}

