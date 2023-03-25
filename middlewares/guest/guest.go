package guest

import (
	"github.com/gofiber/fiber/v2"
)

func GuestMiddleware(c *fiber.Ctx) error {
	if c.Locals("isLogged") == true {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"you must be not logged for access to this resource",
		)
	}
	return c.Next()
}
