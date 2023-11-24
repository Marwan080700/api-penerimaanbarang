package routes

import (
	"pengirimanbarang/handlers"
	"pengirimanbarang/pkg/mysql"
	"pengirimanbarang/repositories"

	"github.com/labstack/echo/v4"
)

func SalesRoutes(e *echo.Group) {
	r := repositories.RepositorySales(mysql.DB)
	h := handlers.HandlerSales(r)

	e.GET("/sales", h.FindSales)
	e.GET("/sale/:id", h.GetSale)
	e.POST("/sale", h.CreateSale)
	e.DELETE("/sale/:id", h.DeleteSale)
	e.PATCH("/sale/:id", h.UpdateSale)
}
