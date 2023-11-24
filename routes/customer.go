package routes

import (
	"pengirimanbarang/handlers"
	"pengirimanbarang/pkg/mysql"
	"pengirimanbarang/repositories"

	"github.com/labstack/echo/v4"
)

func CustomerRoutes(e *echo.Group) {
	r := repositories.RepositoryCustomer(mysql.DB)
	h := handlers.HandlerCustomerCategories(r)

	e.GET("/customers", h.FindCustomer)
	e.GET("/customer/:id", h.GetCustomer)
	e.POST("/customer", h.CreateCustomer)
	e.DELETE("/customer/:id", h.DeleteCustomer)
	e.PATCH("/customer/:id", h.UpdateCustomer)
}
