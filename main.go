package main

import (
	// "os"
	"pengirimanbarang/database"
	"pengirimanbarang/pkg/mysql"
	"pengirimanbarang/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load()

	e := echo.New()
	
	mysql.DatabaseInit()
	database.RunMigration()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))


	routes.RouteInit(e.Group("/api/v1"))

	// e.Static("/uploads", "./uploads")

	// PORT := os.Getenv("PORT")

	// e.Logger.Fatal(e.Start(":" + PORT))
	e.Logger.Fatal(e.Start("localhost:5000"))
}
