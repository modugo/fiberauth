package fiberauth

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

func UserMiddleware(c *fiber.Ctx) error {
	config := c.Locals("auth").(Config)

	c.Locals("isLogged", false)

	tokenString := findAccessToken(c)
	if tokenString != "" {
		token, err := parseAccessToken(c, config, tokenString)
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

func parseAccessToken(c *fiber.Ctx, config Config, tokenString string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.JWTSecret, nil
	})

	// Try to regenerate if cookie-based JWT
	if errors.Is(err, jwt.ErrTokenExpired) && c.Cookies("refreshToken") != "" {
		var login string

		login, err = getLoginFromRefreshToken(c.Cookies("refreshToken"))
		if err != nil {
			return nil, err
		}

		tokenString = generateAccessToken(c, config, login)
		generateRefreshToken(c, config, login)

		return parseAccessToken(c, config, tokenString)
	}

	return token, err
}
