package main

import (
	"log"
	"project/config"
	"project/database"
	"project/middleware"
	"project/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables and config
	config.LoadConfig()

	// Initialize database connection
	database.ConnectDB()

	// Create a new Fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:4000",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Use(middleware.ResponseTimeMiddleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome!")
	})

	// Register routes
	routes.LocationRoutes(app)
	routes.SearchRoutes(app)

	// Start server
	log.Fatal(app.Listen(":4000"))
}
