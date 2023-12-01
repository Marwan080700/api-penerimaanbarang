package routes

import (
	"pengirimanbarang/handlers"
	"pengirimanbarang/pkg/mysql"
	"pengirimanbarang/repositories"

	"github.com/labstack/echo/v4"
)

func ProductCategoriesRoutes(e *echo.Group) {
	r := repositories.RepositoryProductCategories(mysql.DB)
	h := handlers.HandlerProductCategories(r)

	e.GET("/productcategories", h.FindProductCategories)
	e.GET("/productcategories/:id", h.GetProductCategories)
	e.POST("/productcategories", h.CreateProductCategories)
	e.DELETE("/productcategories/:id", h.DeleteProductCategories)
	e.PATCH("/productcategories/:id", h.UpdateProductCategories)
	e.PATCH("/productcategories/cancel/:id", h.CancelProductCategories)
}
