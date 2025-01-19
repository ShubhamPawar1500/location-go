package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Middleware to measure request duration
func ResponseTimeMiddleware(c *fiber.Ctx) error {
	// Start timer
	start := time.Now()

	// Process the request
	err := c.Next()

	// Calculate the duration in nanoseconds
	duration := time.Since(start).Nanoseconds()

	// Store the duration in Locals for later use
	c.Locals("duration", duration)

	return err
}
