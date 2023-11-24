package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	AuthRoutes(e)
	ProductRoutes(e)
	ProductCategoriesRoutes(e)
	CustomerRoutes(e)
	SalesRoutes(e)
	SalesDetailRoutes(e)
	InvoiceRoutes(e)
}
