package main

import (
	"log"
	"os"
	"project/config"
	"project/database"
	"project/middleware"
	"project/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Load environment variables and config
	config.LoadConfig()

	// Initialize database connection
	database.ConnectDB()

	// Create a new Fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("ALLOWED_ORIGIN"),
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

	port := os.Getenv("PORT")

	if port == "" {
		port = "8600"
	}

	// Start server
	log.Fatal(app.Listen(":"+port))
}
