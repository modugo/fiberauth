package fiberauth

import (
	"github.com/modugo/fiberauth/internal/services"
)

type Config struct {
	AccessToken  services.AccessTokenConfiger
	RefreshToken services.RefreshTokenConfiger
	Register     services.RegisterConfiger
	Factors      services.FactorsConfiger
	WithCookie   bool
	WithSignUp   bool
	WithLogout   bool
}

func (c Config) Init() services.Configer {
	if c.AccessToken == nil {
		panic("access token property must be defined in fiberauth.Config")
	}

	if c.RefreshToken == nil {
		c.RefreshToken = RefreshTokenConfig{}
	}

	if c.Register == nil {
		c.Register = RegisterConfig{}
	}

	if c.Factors == nil {
		c.Factors = FactorsConfig{}
	}

	c.Factors = c.Factors.Init()

	return c
}

func (c Config) GetAccessToken() services.AccessTokenConfiger {
	return c.AccessToken
}

func (c Config) GetRefreshToken() services.RefreshTokenConfiger {
	return c.RefreshToken
}

func (c Config) GetRegister() services.RegisterConfiger {
	return c.Register
}

func (c Config) GetFactors() services.FactorsConfiger {
	return c.Factors
}

func (c Config) GetWithCookie() bool {
	return c.WithCookie
}

func (c Config) GetWithSignUp() bool {
	return c.WithSignUp
}

func (c Config) GetWithLogout() bool {
	return c.WithLogout
}
