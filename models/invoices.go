package models

import "time"

type Invoices struct {
	ID                     int       `json:"id" gorm:"primaryKey;autoIncrement"`
	IDSales                int       `json:"-" form:"sale_id" gorm:"type:int"`
	Sales                  Sales     `json:"-" gorm:"foreignKey:IDSales"`
	IDUser                 int       `json:"-" form:"user_id"`
	User                   User      `json:"-" gorm:"foreignKey:IDUser"`
	NumberInvoice          string    `json:"invoice_number" form:"invoice_number" gorm:"type:varchar(255)"`
	DateInvoice            string 	`json:"invoice_date" form:"invoice_date" gorm:"type:varchar(255)"`
	DueDate                string `json:"due_date" form:"due_date" gorm:"type:varchar(255)"`
	SubTotal               int    `json:"sub_total" form:"sub_total" gorm:"type:int"`
	Discount               string    `json:"discount" form:"discount" gorm:"type:varchar(255)"`
	PPN11                  string    `json:"ppn_11" form:"ppn_11" gorm:"type:varchar(255)"`
	TotalAmount            int    `json:"total_amount" form:"total_amount" gorm:"type:int"`
	NoFakturPajak          string    `json:"no_faktur_pajak" form:"no_faktur_pajak" gorm:"type:varchar(255)"`
	NoFakturPajakPengganti string    `json:"no_faktur_pajak_pengganti" form:"no_faktur_pajak_pengganti" gorm:"type:varchar(255)"`
	InvoiceDesc            string    `json:"invoice_description" form:"invoice_description"  gorm:"type:varchar(255)"`
	InvoiceStatus          string    `json:"invoice_status" form:"invoice_status"  gorm:"type:varchar(255)"`
	Approve1               string   `json:"approve_1" form:"approve_1"  gorm:"type:varchar(255)"`
	Approve1Date           string `json:"approve_1_date" form:"approve_1_date"`
	Approve1Desc           string    `json:"approve_1_desc" form:"approve_1_desc" gorm:"type:varchar(255)"`
	Approve2               string   `json:"approve_2" form:"approve_2"  gorm:"type:varchar(255)"`
	Approve2Date           string `json:"approve_2_date" form:"approve_2_date"`
	Approve2Desc           string    `json:"approve_2_desc" form:"approve_2_desc" gorm:"type:varchar(255)"`
	CreatedBy              string    `json:"created_by" gorm:"type:varchar(255)"`
	CreatedAt              time.Time `json:"-"`
	UpdatedBy              string    `json:"updated_by" gorm:"type:varchar(255)"`
	UpdatedAt              time.Time `json:"-"`
}
