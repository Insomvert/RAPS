package Routes

import (
	"golang-fiber/Controllers"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/user/list", Controllers.GetListUser)
	app.Post("/user", Controllers.CreateUser)
}
