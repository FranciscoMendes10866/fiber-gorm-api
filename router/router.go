package router

import (
	"github.com/fiber-gorm-api/handler"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())

	auth := api.Group("/auth")
	auth.Post("/signup", handler.SignUp)
	auth.Post("/signin", handler.SignIn)
}
