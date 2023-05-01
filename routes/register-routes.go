package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pulsarcoder/learn/reactWithgo/controllers"
)

func SetUp(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
}
