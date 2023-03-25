package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/modugo/fiberauth"
	"log"
)

func main() {
	app := fiber.New()

	//fiberauth.Init(app, fiberauth.Config{
	//	WithSignIn:           true,
	//	WithSignUp:           true,
	//	WithLogout:           true,
	//	Entity:               Auth{},
	//	JWTSecret:            []byte("1234"),
	//	AccessTokenDuration:  1 * time.Hour,
	//	RefreshTokenDuration: 24 * time.Hour,
	//	WithRefreshToken:     true,
	//	WithCookie:           true,
	//})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf(
			"I don't need to be logged, but you can get my claims if I'm logged : %v",
			c.Locals("claims"),
		))
	})

	app.Get("/private", fiberauth.AuthMiddleware, func(c *fiber.Ctx) error {
		return c.SendString("I must logged for access to this page!")
	})

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
