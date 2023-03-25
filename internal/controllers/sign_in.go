package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/modugo/fiberauth/internal/models"
	"github.com/modugo/fiberauth/internal/services"
	"github.com/modugo/fiberauth/internal/utils"
)

func SignIn(c *fiber.Ctx) error {
	ctx := c.UserContext()
	cfg := c.Locals("fiberAuth").(services.Configer)

	p := new(models.SignInPayload)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	if !cfg.GetAccessToken().GetRepository().IsValidCredentials(ctx, p.Login, p.Password) {
		return fiber.NewError(fiber.StatusBadRequest, "login or password is incorrect")
	}

	if cfg.GetFactors().GetIsEnabled() && cfg.GetFactors().GetRepository().IsEnabled(ctx, p.Login) {
		return c.JSON(models.TokenResponse{
			Type:         "Factors",
			FactorsToken: utils.GenerateFactorsToken(ctx, c, cfg, p.Login),
		})
	}

	return c.JSON(models.TokenResponse{
		Type:         "JWT",
		AccessToken:  utils.GenerateAccessToken(ctx, c, cfg, p.Login),
		RefreshToken: utils.GenerateRefreshToken(ctx, c, cfg, p.Login),
	})
}
