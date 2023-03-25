package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/modugo/fiberauth/internal/services"
	"github.com/modugo/fiberauth/internal/utils"
	"github.com/pquerna/otp/totp"
	"image/png"
)

type generateTOTPResponse struct {
	SharedSecretKey string `json:"sharedSecretKey"`
	QRCode          string `json:"QRCode"`
}

func GenerateTOTP(c *fiber.Ctx) error {
	cfg := c.Locals("fiberAuth").(services.Configer)

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      cfg.GetFactors().GetTOTP().GetIssuer(),
		AccountName: c.Locals("claims").(jwt.MapClaims)["sub"].(string),
	})
	if err != nil {
		return err
	}

	img, err := key.Image(200, 200)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err = png.Encode(&buf, img); err != nil {
		return err
	}

	return c.JSON(generateTOTPResponse{
		QRCode:          base64.StdEncoding.EncodeToString(buf.Bytes()),
		SharedSecretKey: key.Secret(),
	})
}

type enableTOTPPayload struct {
	SharedSecretKey string `json:"sharedSecretKey"`
	PassCode        string `json:"passCode"`
}

func EnableTOTP(c *fiber.Ctx) error {
	ctx := c.UserContext()
	cfg := c.Locals("fiberAuth").(services.Configer)
	login := c.Locals("claims").(jwt.MapClaims)["sub"].(string)
	repository := cfg.GetFactors().GetTOTP().GetRepository()

	if repository.IsEnabled(ctx, login) {
		return fiber.NewError(fiber.StatusBadRequest, "already enabled")
	}

	p := new(enableTOTPPayload)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	valid := totp.Validate(p.PassCode, p.SharedSecretKey)
	if !valid {
		return fiber.NewError(fiber.StatusBadRequest, "invalid code")
	}

	repository.Enable(ctx, p.SharedSecretKey)

	repository.StoreRecoveryCodes(
		ctx,
		utils.GenerateRecoveryCode(6),
	)

	return nil
}
