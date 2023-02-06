package fiberauth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Entity interface {
	// IsValidCredentials lets know if given credentials are valid
	IsValidCredentials(login, password string) bool

	// IsValidRefreshToken lets know if refresh token is valid
	IsValidRefreshToken(token string) bool

	// GetClaims return data how must be store in JWT token
	GetClaims(login string) jwt.MapClaims

	// DeleteRefreshToken allows to delete refresh token
	DeleteRefreshToken(token string)

	// ManageRefreshToken allows you to perform actions after
	// generating a refresh token (like store it)
	ManageRefreshToken(token string, exp *time.Time)

	/*
	 * THEIR ACTIONS BELOW ARE OPTIONAL BUT U NEED TO INCLUDE IT INTO YOUR ENTITY
	 */

	// IsValidAccessToken is an optional function that allows you
	// to add additional check for validate access token
	IsValidAccessToken(token string) bool

	// ManageAccessToken is an optional function that allows you
	// to perform actions after generating an access token (like store it)
	ManageAccessToken(token string, exp *time.Time)

	// SetCookie is an optional function that allows you to set cookie if
	// you enabled cookie
	SetCookie(name, value string, exp *time.Time) fiber.Cookie
}
