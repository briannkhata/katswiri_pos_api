package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv"
)

func AuthenticateRoutes() fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("apiKey")
		if apiKey == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"msg":     "apiKey header required",
				"data":    nil,
			})
		}

		expectedApiKey := os.Getenv("API_KEY")
		if apiKey != expectedApiKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"msg":     "Invalid API key",
				"data":    nil,
			})
		}

		return c.Next()
	}

}
