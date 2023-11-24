package repositories

import (
	"pengirimanbarang/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindCustomer() ([]models.Customer, error)
	GetCustomer(ID int) (models.Customer, error)
	CreateCustomer(customer models.Customer) (models.Customer, error)
	UpdateCustomer(customer models.Customer) (models.Customer, error)
	DeleteCustomer(customer models.Customer) (models.Customer, error)
}

func RepositoryCustomer(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCustomer() ([]models.Customer, error) {
	var customer []models.Customer
	err := r.db.Find(&customer).Error

	return customer, err
}

func (r *repository) GetCustomer(ID int) (models.Customer, error) {
	var customer models.Customer
	err := r.db.First(&customer, ID).Error

	return customer, err
}

func (r *repository) CreateCustomer(customer models.Customer) (models.Customer, error) {
    err := r.db.Create(&customer).Error

    return customer, err
}



func (r *repository) DeleteCustomer(customer models.Customer) (models.Customer, error) {
	err := r.db.Delete(&customer).Error

	return customer, err
}

func (r *repository) UpdateCustomer(customer models.Customer) (models.Customer, error) {
	err := r.db.Save(&customer).Error

	return customer, err
}
