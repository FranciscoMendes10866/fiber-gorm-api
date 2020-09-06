package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

func GetSponsors(c *fiber.Ctx) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	tokenEmail := claims["email"].(string)
	tokenID := claims["id"].(float64)
	c.JSON(fiber.Map{
		"email": tokenEmail,
		"id":    tokenID,
	})
}
