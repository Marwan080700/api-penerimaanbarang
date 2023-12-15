package repositories

import (
	"pengirimanbarang/models"

	"gorm.io/gorm"
)

type SalesDetailRepository interface {
	FindSalesDetail() ([]models.SalesDetail, error)
	GetSalesDetail(ID int) (models.SalesDetail, error)
	CreateSalesDetail(salesdetail models.SalesDetail) (models.SalesDetail, error)
	UpdateSalesDetail(salesdetail models.SalesDetail) (models.SalesDetail, error)
	DeleteSaleDetail(salesdetail models.SalesDetail) (models.SalesDetail, error)
	GetSalesDetailBySales(IDSales int) ([]models.SalesDetail, error)
}

func RepositorySalesDetail(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindSalesDetail() ([]models.SalesDetail, error) {
	var salesdetail []models.SalesDetail
	err := r.db.Find(&salesdetail).Error

	return salesdetail, err
}

func (r *repository) GetSalesDetail(ID int) (models.SalesDetail, error) {
	var salesdetail models.SalesDetail
	err := r.db.First(&salesdetail, ID).Error

	return salesdetail, err
}

func (r *repository) CreateSalesDetail(salesdetail models.SalesDetail) (models.SalesDetail, error) {
    err := r.db.Create(&salesdetail).Error

    return salesdetail, err
}

func (r *repository) DeleteSaleDetail(salesdetail models.SalesDetail) (models.SalesDetail, error) {
	err := r.db.Delete(&salesdetail).Error

	return salesdetail, err
}

func (r *repository) UpdateSalesDetail(salesdetail models.SalesDetail) (models.SalesDetail, error) {
	err := r.db.Save(&salesdetail).Error

	return salesdetail, err
}

func (r *repository) GetSalesDetailBySales(IDSales int) ([]models.SalesDetail, error) {
    var salesDetail []models.SalesDetail
    err := r.db.Where("id_sales = ?", IDSales).Find(&salesDetail).Error

    return salesDetail, err
}
