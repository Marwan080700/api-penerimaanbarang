package routes

import (
	"pengirimanbarang/handlers"
	"pengirimanbarang/pkg/mysql"
	"pengirimanbarang/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	r := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(r)

	e.GET("/products", h.FindProducts)
	e.GET("/product/:id", h.GetProduct)
	e.POST("/product", h.CreateProduct)
	e.DELETE("/product/:id", h.DeleteProduct)
	e.PATCH("/product/:id", h.UpdateProduct)
	e.GET("/products/category/:id", h.GetProductsByCategoryID)
	
}
