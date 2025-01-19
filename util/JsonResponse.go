package util

import (
	"github.com/gofiber/fiber/v2"
)

func JsonResponse(c *fiber.Ctx, data interface{}) error {
	duration := c.Locals("duration").(int64)
	return c.JSON(fiber.Map{
		"data":     data,
		"duration": duration, // in nanoseconds
	})
}
