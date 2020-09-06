package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/fiber-gorm-api/database"
	"github.com/fiber-gorm-api/model"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) {
	// db connection
	db := database.DBConn
	// user input
	user := new(model.User)
	c.BodyParser(user)
	// hashes the password from the user input
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return
	}
	// exchanges the user input password with the hashed password
	user.Password = string(hash)
	// creates the user
	create := db.Create(user)
	if create == nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}
	// Response
	c.SendStatus(fiber.StatusCreated)
}

func SignIn(c *fiber.Ctx) {
	// dbData response object
	type UserData struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// variable with dbData properties
	var ud UserData
	// db connection
	db := database.DBConn
	// user input
	user := new(model.User)
	c.BodyParser(user)
	// stores the user input password value
	pass := user.Password
	// makes a query to validate if the email exists
	dbData := db.Where(&model.User{Email: user.Email}).Find(user)
	// if the query is null, we will have an error
	// otherwise, we will store every diferent property value to the ud
	if dbData == nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return
	} else {
		ud = UserData{
			ID:       user.ID,
			Email:    user.Email,
			Password: user.Password,
		}
		// here we check if the passwords match
		match := bcrypt.CompareHashAndPassword([]byte(ud.Password), []byte(pass))
		if match != nil {
			c.Status(fiber.StatusUnauthorized)
			return
		}
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = user.Email
		claims["id"] = user.ID
		// Generate encoded token and send it as response.
		loginToken, err := token.SignedString([]byte("FIBERGORM"))
		if err != nil {
			c.SendStatus(fiber.StatusInternalServerError)
			return
		}
		// Response
		c.JSON(fiber.Map{
			"token": loginToken,
		})
	}
}
