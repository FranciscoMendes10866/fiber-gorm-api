package guards

import (
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

func VerifyAuth() func(c *fiber.Ctx) {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
		SigningKey: []byte("FIBERGORM"),
	})
}
