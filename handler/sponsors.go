package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/fiber-gorm-api/database"
	"github.com/fiber-gorm-api/model"
	"github.com/gofiber/fiber"
)

func GetSponsors(c *fiber.Ctx) {
	// selects the filds I want in every single object
	type Sponsor struct {
		ID     int
		Name   string
		Link   string
		Amount int
	}
	// db connection
	db := database.DBConn
	// I'm declaring that I want an array of Sponsor objects
	var sponsors []Sponsor
	// returns all sponsors
	db.Find(&sponsors)
	// reponse
	c.JSON(&sponsors)
}

func CreateSponsor(c *fiber.Ctx) {
	type so struct {
		Name   string
		Link   string
		Amount int
		UserID int
	}
	// db connection
	db := database.DBConn
	sponsors := new(model.Sponsor)
	c.BodyParser(sponsors)
	// token payload
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	tokenID := claims["id"].(float64)
	// converts the tokenID from float64 (1.0000) to int (1)
	var IDtoInt int = int(tokenID)
	// adds the user_id to the object
	sponsors.UserID = IDtoInt
	// creates
	create := db.Create(sponsors)
	c.JSON(create)
}
