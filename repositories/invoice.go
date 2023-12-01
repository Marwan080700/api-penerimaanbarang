package repositories

import (
	"bytes"
	"fmt"
	"pengirimanbarang/models"

	"github.com/jung-kurt/gofpdf"

	"gorm.io/gorm"
)

type InvoiceRepository interface {
	FindInvoices() ([]models.Invoices, error)
	GetInvoice(ID int) (models.Invoices, error)
	CreateInvoice(invoices models.Invoices) (models.Invoices, error)
	UpdateInvoice(invoices models.Invoices) (models.Invoices, error)
	DeleteInvoice(invoices models.Invoices) (models.Invoices, error)
	DeleteInvoiceAndSales(invoiceID int) error
    DeleteSale(sales models.Sales) (models.Sales, error)
	GetSales(ID int) (models.Sales, error)
    CancelInvoice(invoice models.Invoices) (models.Invoices, error)
	PrintInvoice(ID int) (string, error)
}

func RepositoryInvoice(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindInvoices() ([]models.Invoices, error) {
	var invoice []models.Invoices
	err := r.db.Debug().Preload("Sales").Preload("Sales.Customer").Find(&invoice).Error
	if err != nil {
		fmt.Println("Error fetching invoices:", err)
	}

	// Use a different variable name, for example, err2
	err2 := r.db.Debug().Preload("Sales").Preload("Sales.Customer").Find(&invoice).Error

	fmt.Println("Fetched Invoices:", invoice)

	return invoice, err2
}

func (r *repository) GetInvoice(ID int) (models.Invoices, error) {
	var invoice models.Invoices
	err := r.db.Preload("Sales").Preload("Sales.Customer").First(&invoice, ID).Error

	return invoice, err
}

func (r *repository) CreateInvoice(invoices models.Invoices) (models.Invoices, error) {
    err := r.db.Create(&invoices).Error

    return invoices, err
}

func (r *repository) DeleteInvoice(invoices models.Invoices) (models.Invoices, error) {
	err := r.db.Delete(&invoices).Error

	return invoices, err
}

func (r *repository) UpdateInvoice(invoices models.Invoices) (models.Invoices, error) {
	err := r.db.Save(&invoices).Error

	return invoices, err
}

func (r *repository) GetSales(ID int) (models.Sales, error) {
	var sales models.Sales
	err := r.db.First(&sales, ID).Error

	return sales, err
}


func (r *repository) DeleteInvoiceAndSales(invoiceID int) error {
    // Fetch the invoice to get the associated sales ID
    invoice, err := r.GetInvoice(invoiceID)
    if err != nil {
        return err
    }

    // Delete the sales associated with the invoice
    err = r.db.Delete(&invoice.Sales).Error
    if err != nil {
        return err
    }

    // Delete the invoice itself
    err = r.db.Delete(&invoice).Error
    return err
}

// func (r *repository) DeleteSales(sales models.Sales) (models.Sales, error) {
//     err := r.db.Delete(&sales).Error
//     return sales, err
// }

func (r *repository) CancelInvoice(invoice models.Invoices) (models.Invoices, error) {
    // Update the status of the invoice to 1
    invoice.Status = 1

    // Get the associated sales
    sales, err := r.GetSales(invoice.Sales.ID)
    if err != nil {
        return models.Invoices{}, err
    }

    // Update the status of the sales to 1
    sales.Status = 1
    _, err = r.UpdateSale(sales)  // <-- Corrected method name
    if err != nil {
        return models.Invoices{}, err
    }

    // Call the UpdateInvoice method in the repository
    updatedInvoice, err := r.UpdateInvoice(invoice)
    if err != nil {
        return models.Invoices{}, err
    }

    return updatedInvoice, nil
}

func (r *repository) PrintInvoice(ID int) (string, error) {
    // Fetch the necessary data for generating the invoice with preloaded Sales and Customer
    var invoice models.Invoices
    err := r.db.Preload("Sales").Preload("Sales.Customer").Preload("SalesDetail").First(&invoice, ID).Error
    if err != nil {
        return "", fmt.Errorf("failed to fetch invoice data: %v", err)
    }

    // Create a new PDF document
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()

    // Add content to the PDF, you can customize this based on your needs
    pdf.SetFont("Arial", "", 14)
    pdf.Cell(40, 10, fmt.Sprintf("Invoice for Sale ID: %d", invoice.ID))

    // Access the preloaded Sales and Customer data
    if invoice.Sales.Customer.ID != 0 {
        pdf.Cell(40, 10, fmt.Sprintf("Customer Name: %s", invoice.Sales.Customer.NameCustomer))
        // Add more customer details as needed
    }

    // Add more details as needed

    // Save the PDF content to a buffer
    var buf bytes.Buffer
    err = pdf.Output(&buf)
    if err != nil {
        return "", fmt.Errorf("failed to get PDF content: %v", err)
    }

    // Get the PDF content as a string
    content := buf.String()

    return content, nil
}
