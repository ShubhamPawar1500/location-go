package routes

import (
	"project/controllers"

	"github.com/gofiber/fiber/v2"
)

func SearchRoutes(app *fiber.App) {
	app.Get("/search", controllers.GetLocationsByCategoryAndRadius)
	app.Get("/trip-cost/:location_id", controllers.GetTripCost)
}
