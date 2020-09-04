package handler

import (
	"github.com/fiber-gorm-api/database"
	"github.com/fiber-gorm-api/model"
	"github.com/gofiber/fiber"
)

func SignUp(c *fiber.Ctx) {
	db := database.DBConn
	user := new(model.User)
	c.BodyParser(user)
	db.Create(user)
	c.JSON(user)
}

func SignIn(c *fiber.Ctx) {
	db := database.DBConn
	user := new(model.User)
	c.BodyParser(user)
	db.Where(&model.User{Email: user.Email}).Find(user)
	c.JSON(user)
}
