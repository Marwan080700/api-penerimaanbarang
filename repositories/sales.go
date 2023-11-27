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
	DeleteSaleAndAssociatedData(ID int) error  // New method to delete sale and associated data
    GetSalesDetailsBySalesID(salesID int) ([]models.SalesDetail, error)
	CancelSale(sale models.Sales) (models.Sales, error)

    // DeleteSalesDetail deletes a sales detail record
    DeleteSalesDetail(salesDetail models.SalesDetail) (models.SalesDetail, error)
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

func (r *repository) GetSalesDetailsBySalesID(salesID int) ([]models.SalesDetail, error) {
    var salesDetail []models.SalesDetail
    err := r.db.Where("IDSales = ?", salesID).Find(&salesDetail).Error

    return salesDetail, err
}


// DeleteSalesDetail deletes a sales detail record
func (r *repository) DeleteSalesDetail(salesDetail models.SalesDetail) (models.SalesDetail, error) {
    err := r.db.Delete(&salesDetail).Error

    return salesDetail, err
}


func (r *repository) DeleteSaleAndAssociatedData(ID int) error {
	// Get the sale details associated with the sale ID
	salesDetails, err := r.GetSalesDetailsBySalesID(ID)
	if err != nil {
		return err
	}

	// Delete the sales details
	for _, sd := range salesDetails {
		if _, err := r.DeleteSalesDetail(sd); err != nil {
			return err
		}
	}

	// Get the sale by ID
	sale, err := r.GetSale(ID)
	if err != nil {
		return err
	}

	// Delete the sale
	if _, err := r.DeleteSale(sale); err != nil {
		return err
	}

	return nil
}

func (r *repository) CancelSale(sale models.Sales) (models.Sales, error) {
    // Assuming 1 means canceled, update the status
    sale.Status = 1

    // Call the UpdateSale method to persist the changes
    updatedSale, err := r.UpdateSale(sale)
    if err != nil {
        return models.Sales{}, err
    }

    return updatedSale, nil
}
