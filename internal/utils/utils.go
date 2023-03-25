package utils

import (
	"context"
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/modugo/fiberauth/internal/services"
	"math/rand"
	"strings"
	"time"
)

func GenerateAccessToken(ctx context.Context, c *fiber.Ctx, cfg services.Configer, login string) string {
	claims := cfg.GetAccessToken().GetRepository().GetClaims(ctx, login)

	var accessTokenExp *time.Time
	if cfg.GetAccessToken().GetDuration() != 0 {
		t := time.Now().Add(cfg.GetAccessToken().GetDuration())
		accessTokenExp = &t
		claims["exp"] = accessTokenExp.Unix()
	}

	claims["sub"] = login

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessTokenString, err := accessToken.SignedString(cfg.GetAccessToken().GetSecretKey())
	if err != nil {
		panic(err)
	}

	cfg.GetAccessToken().GetRepository().ManageAccessToken(ctx, accessTokenString, accessTokenExp)

	if cfg.GetWithCookie() {
		accessTokenCookie := cfg.GetAccessToken().GetRepository().SetCookie(
			ctx,
			"accessToken",
			accessTokenString,
			accessTokenExp,
		)
		c.Cookie(&accessTokenCookie)
	}

	return accessTokenString
}

func GenerateRefreshToken(ctx context.Context, c *fiber.Ctx, cfg services.Configer, login string) string {
	var refreshTokenExp *time.Time
	if cfg.GetRefreshToken().GetDuration() != 0 {
		t := time.Now().Add(cfg.GetRefreshToken().GetDuration())
		refreshTokenExp = &t
	}

	var refreshToken string
	if cfg.GetRefreshToken().GetIsEnabled() {
		refreshToken = base64.StdEncoding.EncodeToString([]byte(login + "/" + utils.UUIDv4()))
		cfg.GetRefreshToken().GetRepository().ManageRefreshToken(ctx, refreshToken, refreshTokenExp)

		if cfg.GetWithCookie() {
			refreshTokenCookie := cfg.GetAccessToken().GetRepository().SetCookie(
				ctx,
				"refreshToken",
				refreshToken,
				refreshTokenExp,
			)
			c.Cookie(&refreshTokenCookie)
		}
	}

	return refreshToken
}

func GenerateFactorsToken(ctx context.Context, c *fiber.Ctx, cfg services.Configer, login string) string {
	claims := make(jwt.MapClaims)

	var factorsTokenExp *time.Time
	if cfg.GetFactors().GetDuration() != 0 {
		t := time.Now().Add(cfg.GetFactors().GetDuration())
		factorsTokenExp = &t
		claims["exp"] = factorsTokenExp.Unix()
	}

	claims["sub"] = login

	factorsToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	factorsTokenString, err := factorsToken.SignedString(cfg.GetFactors().GetSecretKey())
	if err != nil {
		panic(err)
	}

	if cfg.GetWithCookie() {
		factorsTokenCookie := cfg.GetAccessToken().GetRepository().SetCookie(
			ctx,
			"factorsToken",
			factorsTokenString,
			factorsTokenExp,
		)
		c.Cookie(&factorsTokenCookie)
	}

	return factorsTokenString
}

var InvalidRefreshToken = fiber.NewError(fiber.StatusBadRequest, "invalid refresh token")

func GetLoginFromRefreshToken(token string) (string, error) {
	t, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", InvalidRefreshToken
	}

	split := strings.Split(string(t), "/")
	if len(split) != 2 {
		return "", InvalidRefreshToken
	}

	return split[0], nil
}

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRecoveryCode(nb int) (codes []string) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < nb; i++ {
		code := make([]byte, 10)
		for z := range code {
			code[z] = letterBytes[rand.Intn(len(letterBytes))]
		}
		codes = append(codes, string(code))
	}

	return
}
