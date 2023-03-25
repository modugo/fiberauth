package factors

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/modugo/fiberauth/internal/services"
)

func FactorsMiddleware(c *fiber.Ctx) error {
	var token *jwt.Token
	var tokenString string
	var err error

	noFactorsToken := fiber.NewError(fiber.StatusUnauthorized, "you must provide a factors token")

	cfg := c.Locals("fiberAuth").(services.Configer)

	tokenString = findFactorsToken(c)
	if tokenString == "" {
		return noFactorsToken
	}

	token, err = parseFactorsToken(cfg, tokenString)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || err != nil {
		return noFactorsToken
	}

	c.Locals("factorsClaims", claims)

	return c.Next()
}

func findFactorsToken(c *fiber.Ctx) string {
	// With cookie
	if c.Cookies("factorsToken") != "" {
		return c.Cookies("factorsToken")
	}

	// With header
	return c.Get("factors-token")
}

func parseFactorsToken(cfg services.Configer, tokenString string) (token *jwt.Token, err error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return cfg.GetFactors().GetSecretKey(), nil
	})
}
