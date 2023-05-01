package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pulsarcoder/learn/reactWithgo/database"
	"github.com/pulsarcoder/learn/reactWithgo/routes"
)

func main() {
	// connect, err := gorm.Open(sqlite.Open("finData.db"), &gorm.Config{})

	// if err != nil {
	// 	panic("Database not connected properly")
	// }

	// connect.AutoMigrate()
	database.ConnectDatabase()

	app := fiber.New()

	routes.SetUp(app)

	app.Listen(":8081")

}
