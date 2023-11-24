package routes

import (
	"pengirimanbarang/handlers"
	"pengirimanbarang/pkg/mysql"
	"pengirimanbarang/repositories"

	"github.com/labstack/echo/v4"
)

func SalesDetailRoutes(e *echo.Group) {
	r := repositories.RepositorySalesDetail(mysql.DB)
	h := handlers.HandlerSalesDetail(r)

	e.GET("/sales-details", h.FindSalesDetail)
	e.GET("/sales-detail/:id", h.GetSalesDetail)
	e.POST("/sales-detail", h.CreateSalesDetail)
	e.DELETE("/sales-detail/:id", h.DeleteSalesDetail)
	e.PATCH("/sales-detail/:id", h.UpdateSalesDetail)
}
