package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Auth struct{}

func (Auth) IsValidCredentials(login, password string) bool {
	return login == "admin" && password == "admin"
}

func (Auth) IsValidRefreshToken(token string) bool {
	return true
}

func (Auth) GetClaims(login string) jwt.MapClaims {
	return jwt.MapClaims{
		"jti":    utils.UUIDv4(),
		"sub":    "1",
		"scopes": []string{"member", "admin"},
	}
}

func (Auth) SetCookie(name, value string, exp *time.Time) fiber.Cookie {
	c := fiber.Cookie{
		Name:  name,
		Value: value,
	}

	if exp != nil {
		c.Expires = *exp
	}

	return c
}

func (Auth) DeleteRefreshToken(token string) {}

func (Auth) ManageAccessToken(token string, exp *time.Time) {}

func (Auth) ManageRefreshToken(token string, exp *time.Time) {}

func (Auth) IsValidAccessToken(token string) bool {
	return true
}
