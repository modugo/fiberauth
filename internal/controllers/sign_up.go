package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/modugo/fiberauth/internal/services"
)

func SignUp(c *fiber.Ctx) error {
	ctx := c.UserContext()
	cfg := c.Locals("fiberAuth").(services.Configer)

	var p map[string]interface{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	result, err := cfg.GetRegister().GetRepository().CreateAccount(ctx, p)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
