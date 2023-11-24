package database

import (
	"fmt"
	"pengirimanbarang/models"
	"pengirimanbarang/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Customer{},
		&models.Product{},
		&models.ProductCategories{},
		&models.Sales{},
		&models.SalesDetail{},
		&models.Invoices{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration Succes")
}
