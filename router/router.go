package router

import (
	"github.com/fiber-gorm-api/guards"
	"github.com/fiber-gorm-api/handler"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())

	auth := api.Group("/auth")
	auth.Post("/signup", handler.SignUp)
	auth.Post("/signin", handler.SignIn)

	sponsors := api.Group("/sponsors")
	sponsors.Get("/", guards.VerifyAuth(), handler.GetSponsors)
	sponsors.Get("/user", guards.VerifyAuth(), handler.GetUserSponsors)
	sponsors.Post("/", guards.VerifyAuth(), handler.CreateSponsor)
	sponsors.Get("/total", guards.VerifyAuth(), handler.CountTotalSponsors)
}
