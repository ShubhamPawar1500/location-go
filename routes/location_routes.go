package routes

import (
	"project/controllers"

	"github.com/gofiber/fiber/v2"
)

func LocationRoutes(app *fiber.App) {
	locationGroup := app.Group("/locations")
	locationGroup.Post("/", controllers.CreateLocation)
	locationGroup.Get("/:category", controllers.GetLocationByCategory)
}
