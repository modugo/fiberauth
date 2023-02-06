# fiberauth
> ⚡️🔒 Plug and play JWT authentication module for Fiber

## ⚙️ Installation
```shell
go get -u github.com/modugo/fiberauth
```

## ⚡️ Quickstart
```go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/modugo/fiberauth"
)

func main() {
	app := fiber.New()

	fiberauth.Init(app, fiberauth.Config{
		WithSignIn:           true,
		WithSignUp:           true,
		WithRefresh:          true,
		WithLogout:           true,
		Entity:               Auth{},
		JWTSecret:            []byte("1234"),
		AccessTokenDuration:  1 * time.Hour,
		RefreshTokenDuration: 24 * time.Hour,
		WithRefreshToken:     true,
		WithCookie:           true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8080")
}
```

## ⚙️ Configuration
You can configure `fiberauth` easily with the `fiberauth.Config`,
see below for know how to configure it :




| Name                 | Type            | Usage                                               |
|----------------------|-----------------|-----------------------------------------------------|
| WithSignIn           | `bool`          | Enable/Disable Sign-In endpoint                     |
| WithSignUp           | `bool`          | Enable/Disable Sign-Up endpoint                     |
| WithRefresh          | `bool`          | Enable/Disable Refresh endpoint                     |
| WithLogout           | `bool`          | Enable/Disable Logout endpoint                      |
| Entity               | `interface`     | A struct contains all methods in `Entity` interface |
| JWTSecret            | `[]byte`        | Private secret key for JWT                          |
| AccessTokenDuration  | `time.Duration` | Define access token duration                        |
| RefreshTokenDuration | `time.Duration` | Define refresh token duration                       |
| WithRefreshToken     | `bool`          | Enable/Disable refresh token feature                |
| WithCookie           | `bool`          | Enable/Disable cookie-based JWT                     |