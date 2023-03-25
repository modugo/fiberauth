package fiberauth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/modugo/fiberauth/internal/controllers"
	"github.com/modugo/fiberauth/internal/middlewares/factors"
	"github.com/modugo/fiberauth/internal/services"
	auth2 "github.com/modugo/fiberauth/middlewares/auth"
	"github.com/modugo/fiberauth/middlewares/guest"
	"github.com/modugo/fiberauth/middlewares/user"
)

func Init(app *fiber.App, cfg services.Configer, handlers ...fiber.Handler) {
	// Init default value
	cfg = cfg.Init()

	// Global middleware for access to auth config
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("fiberAuth", cfg)
		return c.Next()
	})

	// Global middleware for parse jwt token if exists
	app.Use(user.UserMiddleware)

	auth := app.Group("/auth", handlers...)

	auth.Post("/sign-in", controllers.SignIn)

	if cfg.GetRegister().GetIsEnabled() {
		auth.Post("/sign-up", guest.GuestMiddleware, controllers.SignUp)
	}

	if cfg.GetRefreshToken().GetIsEnabled() {
		auth.Post("/refresh", controllers.Refresh)
	}

	if cfg.GetWithLogout() {
		auth.Post("/logout", controllers.Logout)
	}

	f := app.Group("/factors")

	f.Get("/", factors.FactorsMiddleware, controllers.GetAvailableFactors)
	f.Post("/", factors.FactorsMiddleware, controllers.UseFactors)

	if cfg.GetFactors().GetTOTP().GetIsEnabled() {
		totp := f.Group("/totp", auth2.AuthMiddleware)
		totp.Post("/generate", controllers.GenerateTOTP)
		totp.Post("/enable", controllers.EnableTOTP)
		//totp.Delete("/disable")
	}

	if cfg.GetFactors().GetSMS().GetIsEnabled() {

	}
}
