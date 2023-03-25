package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/modugo/fiberauth/internal/services"
	"time"
)

func Logout(c *fiber.Ctx) error {
	ctx := c.UserContext()
	cfg := c.Locals("fiberAuth").(services.Configer)

	if cfg.GetWithCookie() {
		t := time.Now().AddDate(-1, 0, 0)

		accessTokenCookie := cfg.GetAccessToken().GetRepository().SetCookie(ctx, "accessToken", "", &t)
		c.Cookie(&accessTokenCookie)

		if cfg.GetRefreshToken().GetIsEnabled() {
			if c.Cookies("refreshToken") != "" {
				cfg.GetRefreshToken().GetRepository().DeleteRefreshToken(ctx, c.Cookies("refreshToken"))
			}
			refreshTokenCookie := cfg.GetAccessToken().GetRepository().SetCookie(ctx, "refreshToken", "", &t)
			c.Cookie(&refreshTokenCookie)
		}
	}

	return nil
}
