package handler

import (
	"github.com/fiber-gorm-api/database"
	"github.com/fiber-gorm-api/model"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func SignUp(c *fiber.Ctx) {
	db := database.DBConn
	user := new(model.User)
	c.BodyParser(user)
	hash, _ := HashPassword(user.Password)
	user.Password = hash
	db.Create(user)
	c.SendStatus(fiber.StatusCreated)
}

func SignIn(c *fiber.Ctx) {
	db := database.DBConn
	user := new(model.User)
	c.BodyParser(user)
	db.Where(&model.User{Email: user.Email}).Find(user)
	c.JSON(user)
}
