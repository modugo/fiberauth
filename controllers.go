package fiberauth

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type TokenResponse struct {
	Type         string `json:"type"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

type SignInPayload struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func signIn(c *fiber.Ctx) error {
	config := c.Locals("auth").(Config)

	p := new(SignInPayload)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	if !config.Entity.IsValidCredentials(p.Login, p.Password) {
		return fiber.NewError(fiber.StatusBadRequest, "login or password is incorrect")
	}

	return c.JSON(TokenResponse{
		Type:         "JWT",
		AccessToken:  generateAccessToken(c, config, p.Login),
		RefreshToken: generateRefreshToken(c, config, p.Login),
	})
}

func signUp(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

type RefreshPayload struct {
	Token string `json:"token"`
}

func refresh(c *fiber.Ctx) error {
	config := c.Locals("auth").(Config)

	p := new(RefreshPayload)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	if !config.Entity.IsValidRefreshToken(p.Token) {
		return fiber.NewError(fiber.StatusBadRequest, "refresh token not valid")
	}

	login, err := getLoginFromRefreshToken(p.Token)
	if err != nil {
		return err
	}

	return c.JSON(TokenResponse{
		Type:         "JWT",
		AccessToken:  generateAccessToken(c, config, login),
		RefreshToken: generateRefreshToken(c, config, login),
	})
}

func logout(c *fiber.Ctx) error {
	config := c.Locals("auth").(Config)

	if config.WithCookie {
		t := time.Now().AddDate(-1, 0, 0)

		accessTokenCookie := config.Entity.SetCookie("accessToken", "", &t)
		c.Cookie(&accessTokenCookie)

		if config.WithRefreshToken {
			if c.Cookies("refreshToken") != "" {
				config.Entity.DeleteRefreshToken(c.Cookies("refreshToken"))
			}
			refreshTokenCookie := config.Entity.SetCookie("refreshToken", "", &t)
			c.Cookie(&refreshTokenCookie)
		}
	}

	return nil
}
