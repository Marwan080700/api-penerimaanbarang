package routes

import (
	"pengirimanbarang/handlers"
	"pengirimanbarang/pkg/mysql"
	"pengirimanbarang/repositories"

	"github.com/labstack/echo/v4"
)

func InvoiceRoutes(e *echo.Group) {
    r := repositories.RepositoryInvoice(mysql.DB)
    h := handlers.HandlerInvoice(r)

    // Existing routes...
    e.GET("/invoices", h.FindInvoices)
    e.GET("/invoice/:id", h.GetInvoice)
    e.POST("/invoice", h.CreateInvoice)
    e.DELETE("/invoice/:id", h.DeleteInvoice)
    e.PATCH("/invoice/:id", h.UpdateInvoice)
    e.PATCH("/invoice/cancel/:id", h.CancelInvoice)
    e.GET("/invoices/print/:id", h.PrintInvoice)
}