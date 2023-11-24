package repositories

import (
	"pengirimanbarang/models"

	"gorm.io/gorm"
)

type SalesRepository interface {
	FindSales() ([]models.Sales, error)
	GetSale(ID int) (models.Sales, error)
	CreateSale(sales models.Sales) (models.Sales, error)
	UpdateSale(sales models.Sales) (models.Sales, error)
	DeleteSale(sales models.Sales) (models.Sales, error)
}

func RepositorySales(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindSales() ([]models.Sales, error) {
	var sales []models.Sales
	err := r.db.Find(&sales).Error

	return sales, err
}

func (r *repository) GetSale(ID int) (models.Sales, error) {
	var sales models.Sales
	err := r.db.First(&sales, ID).Error

	return sales, err
}

func (r *repository) CreateSale(sales models.Sales) (models.Sales, error) {
    err := r.db.Create(&sales).Error

    return sales, err
}

func (r *repository) DeleteSale(sales models.Sales) (models.Sales, error) {
	err := r.db.Delete(&sales).Error

	return sales, err
}

func (r *repository) UpdateSale(sales models.Sales) (models.Sales, error) {
	err := r.db.Save(&sales).Error

	return sales, err
}
