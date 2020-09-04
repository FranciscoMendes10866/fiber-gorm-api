package handler

import (
	"github.com/fiber-gorm-api/database"
	"github.com/fiber-gorm-api/model"
	"github.com/gofiber/fiber"
)

func SignUp(c *fiber.Ctx) {
	db := database.DBConn
	body := new(model.User)
	c.BodyParser(body)
	user := db.Create(body)
	c.JSON(user)
}

func SignIn(c *fiber.Ctx) {
	db := database.DBConn
	body := new(model.User)
	c.BodyParser(body)
	user := db.Find(body)
	c.JSON(user)
}
