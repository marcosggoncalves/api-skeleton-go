package server

import (
	"ApiSup/internal/database"
	"ApiSup/internal/middlewares"
	"ApiSup/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.ConfigureCORS())

	database.Initialize()

	routes.InitializeRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
