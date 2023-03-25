package auth

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	if c.Locals("isLogged") == false {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"you are not authenticated",
		)
	}
	return c.Next()
}
