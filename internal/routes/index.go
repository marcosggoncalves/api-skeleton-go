package routes

import (
	controller "ApiSup/internal/controllers"
	"ApiSup/internal/middlewares"
	"ApiSup/internal/repository"
	"ApiSup/internal/services"

	"github.com/labstack/echo/v4"
)

func InitializeRoutes(e *echo.Echo) {
	userRepo := repository.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	auth := e.Group("/auth")
	auth.POST("/login", userController.Login)

	protected := e.Group("/api")
	protected.Use(middlewares.VerifyTokenHandler())

	protected.POST("/usuarios", userController.CreateUser)
	protected.GET("/usuarios", userController.GetUsers)
	protected.GET("/usuarios/:id", userController.GetUser)
	protected.PUT("/usuarios/:id", userController.UpdateUser)
	protected.DELETE("/usuarios/:id", userController.DeleteUser)
}
