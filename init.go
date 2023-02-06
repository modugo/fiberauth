package fiberauth

import (
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App, config Config, handlers ...fiber.Handler) {
	// Global middleware for access to auth config
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("auth", config)
		return c.Next()
	})

	// Global middleware for parse jwt token if exists
	app.Use(UserMiddleware)

	auth := app.Group("/auth", handlers...)

	if config.WithSignIn {
		auth.Post("/sign-in", signIn)
	}

	if config.WithSignUp {
		auth.Post("/sign-up", signUp)
	}

	if config.WithRefresh {
		auth.Post("/refresh", refresh)
	}

	if config.WithLogout {
		auth.Post("/logout", logout)
	}
}
