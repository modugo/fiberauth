package user

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/modugo/fiberauth/internal/services"
	"github.com/modugo/fiberauth/internal/utils"
	"strings"
)

func UserMiddleware(c *fiber.Ctx) error {
	var token *jwt.Token
	var tokenString string
	var err error

	cfg := c.Locals("fiberAuth").(services.Configer)

	c.Locals("isLogged", false)

	tokenString = findAccessToken(c)

	if tokenString == "" && c.Cookies("refreshToken") != "" {
		tokenString, err = useRefreshToken(c, cfg, c.Cookies("refreshToken"))
	}

	if tokenString != "" && err == nil {
		token, err = parseAccessToken(c, cfg, tokenString)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && err == nil {
			c.Locals("claims", claims)
			c.Locals("isLogged", true)
		}
	}

	return c.Next()
}

func findAccessToken(c *fiber.Ctx) string {
	// With cookie
	if c.Cookies("accessToken") != "" {
		return c.Cookies("accessToken")
	}

	// With header
	if c.Get("authorization") != "" {
		components := strings.SplitN(c.Get("authorization"), " ", 2)
		if len(components) == 2 && components[0] == "Bearer" {
			return components[1]
		}
	}

	return ""
}

func parseAccessToken(c *fiber.Ctx, cfg services.Configer, tokenString string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return cfg.GetAccessToken().GetSecretKey(), nil
	})

	// Try to regenerate if cookie-based JWT
	if errors.Is(err, jwt.ErrTokenExpired) && c.Cookies("refreshToken") != "" {
		tokenString, err = useRefreshToken(c, cfg, c.Cookies("refreshToken"))
		if err != nil {
			return
		}

		return parseAccessToken(c, cfg, tokenString)
	}

	return
}

func useRefreshToken(c *fiber.Ctx, cfg services.Configer, refreshToken string) (token string, err error) {
	ctx := c.UserContext()

	var login string

	login, err = utils.GetLoginFromRefreshToken(refreshToken)
	if err != nil {
		return
	}

	token = utils.GenerateAccessToken(ctx, c, cfg, login)
	utils.GenerateRefreshToken(ctx, c, cfg, login)

	return
}
