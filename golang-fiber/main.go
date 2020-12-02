package main

import (
	"fmt"
	"golang-fiber/Config"
	"golang-fiber/Models"
	"golang-fiber/Routes"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDbConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.User{})

	app := fiber.New()

	Routes.SetupRoutes(app)

	app.Listen(3000)
}
