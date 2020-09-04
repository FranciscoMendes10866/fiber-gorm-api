package main

import (
	"fmt"

	"github.com/fiber-gorm-api/model"

	"github.com/fiber-gorm-api/database"
	"github.com/fiber-gorm-api/router"
	"github.com/gofiber/fiber"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() {
	var err error
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=8000 sslmode=disable"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database!🔥")
	database.DBConn.AutoMigrate(&model.User{}, &model.Sponsor{})
	fmt.Println("Database Migrated!🧨")
}

func main() {
	app := fiber.New()

	initDB()

	router.SetupRoutes(app)

	app.Listen(3030)
}
