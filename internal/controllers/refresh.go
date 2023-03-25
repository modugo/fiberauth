package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/modugo/fiberauth/internal/models"
	"github.com/modugo/fiberauth/internal/services"
	"github.com/modugo/fiberauth/internal/utils"
)

func Refresh(c *fiber.Ctx) error {
	ctx := c.UserContext()
	cfg := c.Locals("fiberAuth").(services.Configer)

	p := new(models.RefreshPayload)
	if err := c.BodyParser(p); err != nil && err != fiber.ErrUnprocessableEntity {
		return err
	}

	if p.Token == "" {
		p.Token = c.Cookies("refreshToken")
		if p.Token == "" {
			return fiber.NewError(fiber.StatusBadRequest, "refresh token not valid")
		}
	}

	if !cfg.GetRefreshToken().GetRepository().IsValidRefreshToken(ctx, p.Token) {
		return fiber.NewError(fiber.StatusBadRequest, "refresh token not valid")
	}

	login, err := utils.GetLoginFromRefreshToken(p.Token)
	if err != nil {
		return err
	}

	cfg.GetRefreshToken().GetRepository().DeleteRefreshToken(ctx, p.Token)

	return c.JSON(models.TokenResponse{
		Type:         "JWT",
		AccessToken:  utils.GenerateAccessToken(ctx, c, cfg, login),
		RefreshToken: utils.GenerateRefreshToken(ctx, c, cfg, login),
	})
}
