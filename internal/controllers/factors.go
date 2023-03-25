package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/modugo/fiberauth/internal/models"
	"github.com/modugo/fiberauth/internal/services"
	"github.com/modugo/fiberauth/internal/utils"
	"github.com/pquerna/otp/totp"
	"strings"
	"time"
)

type availableFactorsResponse struct {
	Default   string   `json:"default"`
	Available []string `json:"available"`
}

func GetAvailableFactors(c *fiber.Ctx) error {
	ctx := c.UserContext()
	cfg := c.Locals("fiberAuth").(services.Configer)
	login := c.Locals("factorsClaims").(jwt.MapClaims)["sub"].(string)

	return c.JSON(availableFactorsResponse{
		Default:   cfg.GetFactors().GetRepository().GetDefaultFactor(ctx, login),
		Available: cfg.GetFactors().GetRepository().GetAvailableFactors(ctx, login),
	})
}

type useFactorsPayload struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

func UseFactors(c *fiber.Ctx) error {
	ctx := c.UserContext()
	cfg := c.Locals("fiberAuth").(services.Configer)
	login := c.Locals("factorsClaims").(jwt.MapClaims)["sub"].(string)

	p := new(useFactorsPayload)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	switch strings.ToLower(p.Type) {
	case "totp":
		valid := totp.Validate(
			p.Code,
			cfg.GetFactors().GetTOTP().GetRepository().GetSharedSecretKey(ctx, login),
		)
		if !valid {
			return fiber.NewError(fiber.StatusBadRequest, "invalid code")
		}
	case "sms":
	default:
		return fiber.NewError(fiber.StatusBadRequest, "given factor type not supported")
	}

	if cfg.GetWithCookie() {
		t := time.Now().AddDate(-1, 0, 0)
		factorsTokenCookie := cfg.GetAccessToken().GetRepository().SetCookie(ctx, "factorsToken", "", &t)
		c.Cookie(&factorsTokenCookie)
	}

	return c.JSON(models.TokenResponse{
		Type:         "JWT",
		AccessToken:  utils.GenerateAccessToken(ctx, c, cfg, login),
		RefreshToken: utils.GenerateRefreshToken(ctx, c, cfg, login),
	})
}
