package fiberauth

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

func generateAccessToken(c *fiber.Ctx, config Config, login string) string {
	claims := config.Entity.GetClaims(login)

	var accessTokenExp *time.Time
	if config.AccessTokenDuration != 0 {
		t := time.Now().Add(config.AccessTokenDuration)
		accessTokenExp = &t
		claims["exp"] = accessTokenExp.Unix()
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessTokenString, err := accessToken.SignedString(config.JWTSecret)
	if err != nil {
		panic(err)
	}

	config.Entity.ManageAccessToken(accessTokenString, accessTokenExp)

	if config.WithCookie {
		accessTokenCookie := config.Entity.SetCookie("accessToken", accessTokenString, accessTokenExp)
		c.Cookie(&accessTokenCookie)
	}

	return accessTokenString
}

func generateRefreshToken(c *fiber.Ctx, config Config, login string) string {
	var refreshTokenExp *time.Time
	if config.RefreshTokenDuration != 0 {
		t := time.Now().Add(config.RefreshTokenDuration)
		refreshTokenExp = &t
	}

	var refreshToken string
	if config.WithRefreshToken {
		refreshToken = base64.StdEncoding.EncodeToString([]byte(login + "." + utils.UUIDv4()))
		config.Entity.ManageRefreshToken(refreshToken, refreshTokenExp)

		if config.WithCookie {
			refreshTokenCookie := config.Entity.SetCookie("refreshToken", refreshToken, refreshTokenExp)
			c.Cookie(&refreshTokenCookie)
		}
	}

	return refreshToken
}

var InvalidRefreshToken = fiber.NewError(fiber.StatusBadRequest, "invalid refresh token")

func getLoginFromRefreshToken(token string) (string, error) {
	t, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", InvalidRefreshToken
	}

	split := strings.Split(string(t), ".")
	if len(split) != 2 {
		return "", InvalidRefreshToken
	}

	return split[0], nil
}
