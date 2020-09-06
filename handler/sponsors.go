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

func CountTotalSponsors(c *fiber.Ctx) {
	var count int64
	db := database.DBConn
	sponsors := new(model.Sponsor)
	c.BodyParser(sponsors)
	db.Model(sponsors).Count(&count)
	c.JSON(fiber.Map{
		"total": &count,
	})
}

func GetUserSponsors(c *fiber.Ctx) {
	// selects the filds I want in every single object
	type Sponsor struct {
		ID     int
		Name   string
		Link   string
		Amount int
		UserID int
	}
	db := database.DBConn
	var sponsors []Sponsor
	// token payload
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	tokenID := claims["id"].(float64)
	var IDtoInt int = int(tokenID)
	// query
	db.Where("user_id = ?", IDtoInt).Find(&sponsors)
	c.JSON(&sponsors)
}

func DeleteSingleSponsor(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var sponsor model.Sponsor
	db.First(&sponsor, id)
	if sponsor.Name == "" {
		c.JSON(fiber.Map{"err": "An error occored."})
		return
	}
	db.Delete(&sponsor)
	c.JSON(fiber.Map{
		"msg": "Was deleted.",
	})
}
