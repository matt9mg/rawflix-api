package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matt9mg/rawflix-api/repositories"
	"github.com/matt9mg/rawflix-api/services"
)

type LogoutController interface {
	Logout(ctx *fiber.Ctx) error
}

type Logout struct {
	jwt      services.JWTAuthenticator
	userRepo repositories.UserRepository
}

func NewLogout(jwt services.JWTAuthenticator, userRepo repositories.UserRepository) LogoutController {
	return &Logout{
		jwt:      jwt,
		userRepo: userRepo,
	}
}

func (l *Logout) Logout(ctx *fiber.Ctx) error {
	claims, err := l.jwt.Validate(ctx)

	if err != nil {
		return ctx.JSON("OK")
	}

	_ = l.userRepo.RemoveTokenFromByUserID(claims.UserID)

	return ctx.JSON("OK")
}
